#!/usr/bin/env ruby

Command = Struct.new('Command', :file, :arg, :output)
Directory = Struct.new('Directory', :name, :parent, :children)
FileObject = Struct.new('FileObject', :name, :parent, :size)

def parse_commands
  lines = ARGF.readlines
  commands = Array.new

  lines.each do |line|
    tokens = line.split
    if tokens.first == '$'
      commands.push(Command.new(tokens[1], tokens[2]))
    else
      commands.last.output ||= Array.new
      commands.last.output.push(line.chomp.split)
    end
  end

  commands
end

def parse_fs(commands)
  current_dir = root = Directory.new('/')

  commands.each do |c|
    case c.file
    when 'cd'
      if c.arg == '/'
        current_dir = root
      elsif c.arg == '..'
        raise 'cd .., current directory has no parent' if current_dir.parent.nil?
        current_dir = current_dir.parent
      else
        raise 'Current Dir has no children' if current_dir.children.nil?
        current_dir = current_dir.children.select { |x| x.name == c.arg }.first
        raise "cd #{c.arg} is not a directory" unless current_dir.is_a?(Directory)
      end
    when 'ls'
      c.output.each do |o|
        child = nil
        if o[0] == 'dir'
          child = Directory.new(o[1], current_dir)
        else
          child = FileObject.new(o[1], current_dir, o[0].to_i)
        end
        current_dir.children ||= Array.new
        current_dir.children.push(child)
      end
    else
      raise 'uh oh'
    end
  end

  root
end

def calc_entry_size(entry, sizes)
  if entry.is_a?(FileObject)
    return entry.size
  else
    size = entry.children.inject(0) { |sum, e| sum += calc_entry_size(e, sizes) }
    sizes.push(size)
    return size
  end
end

def calc_dir_sizes(root)
  sizes = Array.new
  calc_entry_size(root, sizes)
  sizes
end

commands = parse_commands
root = parse_fs(commands)
sizes = calc_dir_sizes(root)
answer = sizes.inject(0) { |sum, s| sum = s < 100000 ? sum + s : sum }
puts answer
