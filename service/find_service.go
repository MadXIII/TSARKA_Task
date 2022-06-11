package service

type Finder struct{}

func NewFinder() *Finder {
	return &Finder{}
}

func (s *Finder) CheckStr(str string) string {
	if len(str) == 1 {
		return str
	}

	return FindSubstring(str)
}

func FindSubstring(str string) string {
	have := map[byte]struct{}{} // TODO: change to map[byte]struct{}
	position := map[byte]int{}
	var max int
	var prePosition int
	var start int
	var end int
	var i int

	for i = 0; i < len(str); i++ {
		if _, ok := have[str[i]]; !ok {
			have[str[i]] = struct{}{}
			position[str[i]] = i
			continue
		}
		if i-prePosition+1 > max {
			start = prePosition
			end = i
			max = i - prePosition + 1
		}
		prePosition = position[str[i]] + 1
		i = position[str[i]]
		have = map[byte]struct{}{}
		position = map[byte]int{}
	}

	if i-prePosition+1 > max {
		start = prePosition
		end = i
	}
	return str[start:end]
}
