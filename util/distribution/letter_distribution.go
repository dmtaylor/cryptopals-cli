package distribution

import "sort"

type LetterDistribution struct {
	dist map[rune]int
}

func NewLetterDistribution() *LetterDistribution {
	return &LetterDistribution{
		dist: make(map[rune]int),
	}
}

func (l *LetterDistribution) Add(r rune, n int) error {
	if _, ok := l.dist[r]; ok {
		l.dist[r] += n
	} else {
		l.dist[r] = n
	}
	return nil
}

func (l *LetterDistribution) Ordering() []rune {
	keys := make([]rune, 0, len(l.dist))
	for key := range l.dist {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return l.dist[keys[i]] > l.dist[keys[j]]
	})

	return keys
}

func (l *LetterDistribution) Clear() {
	l.dist = make(map[rune]int)
}
