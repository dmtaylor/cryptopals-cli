package distribution

import "sort"

type ByteDistribution struct {
	dist map[byte]int
}

func NewByteDistribution() *ByteDistribution {
	return &ByteDistribution{
		dist: make(map[byte]int),
	}
}

func (d *ByteDistribution) Add(b byte, n int) error {
	if _, ok := d.dist[b]; ok {
		d.dist[b] += n
	} else {
		d.dist[b] = n
	}
	return nil
}

func (d *ByteDistribution) Ordering() []byte {
	keys := make([]byte, 0, len(d.dist))
	for key := range d.dist {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return d.dist[keys[i]] > d.dist[keys[j]]
	})
	return keys
}

func (d *ByteDistribution) Clear() {
	d.dist = make(map[byte]int)
}
