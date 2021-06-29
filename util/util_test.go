package util

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpSplit(t *testing.T) {
	result := RegexpSplit(regexp.MustCompile(`\p{Han}+`),
		"BP神经网络如何训练才能在分类时增加区分度？", -1)
	if len(result) != 2 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`(\p{Han})+`),
		"BP神经网络如何训练才能在分类时增加区分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`([\p{Han}#]+)`),
		",BP神经网络如何训练才能在分类时#增加区分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
}

func TestLoadProb(t *testing.T) {
	start := map[byte]float64{
		'B': -1.6134773264811813, 'E': -3.14e+100, 'M': -3.14e+100, 'S': -0.22213624217341046,
	}
	trans := map[byte]map[byte]float64{
		'B': {'E': -0.08019685120220049, 'M': -2.5631014860381245},
		'E': {'B': -1.9211369858872491, 'S': -0.15833987018971005},
		'M': {'E': -0.23021182203175303, 'M': -1.5816540855242822},
		'S': {'B': -2.1198043739881163, 'S': -0.12789600087693714},
	}
	assert.Equal(t, start, LoadProbStartFromJsonFile("../id_prob_start.py"))
	assert.Equal(t, trans, LoadProbEmitFromJsonFile("../id_prob_trans.py"))
}
