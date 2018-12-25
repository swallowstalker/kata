
def meeting(s)
  s.upcase!

  name_mapping = {}

  names = s.split(';')
  names.each do |full_name|
    name_fraction = full_name.split(':')
    name_mapping[name_fraction[1]] = [] unless name_mapping.key?(name_fraction[1])
    name_mapping[name_fraction[1]].push(name_fraction[0])
  end

  sorted_name_mapping = {}
  name_mapping.keys.sort.each do |k|
    name_mapping[k].sort!
    sorted_name_mapping[k] = name_mapping[k]
  end

  format(sorted_name_mapping)
end

def format(name_mapping)
  formatted_names = ''
  name_mapping.each do |last_name, front_names|
    front_names.each do |front_name|
      formatted_names += "(#{last_name}, #{front_name})"
    end
  end

  formatted_names
end

meeting('Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Megan:Stan;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Alex:Korn')
meeting('Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill')