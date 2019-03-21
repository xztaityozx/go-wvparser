package wvparser

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestWVParser_GetDocument(t *testing.T) {
	p := filepath.Join("test", "target.csv")
	wp := WVParser{FilePath: p}
	actual, err := wp.GetDocument()
	assert.NoError(t, err)
	expect := Document{
		Header: Header{
			Format:    "2dsweep MONTE_CARLO",
			Signal:    "m8d",
			SavedTime: "15:12:20 Mon Feb 25 2019",
		},
		TextElements: []TextElement{
			`sweep 1
 2.5000E-09 , 8.0001E-01 
 1.0000E-08 , 1.0535E-04 
 1.7500E-08 , 1.6397E-04 
`,
			`sweep 2
 2.5000E-09 , 8.0000E-01 
 1.0000E-08 , 1.9597E-04 
 1.7500E-08 , 2.6414E-04 
`,
			`sweep 3
 2.5000E-09 , 7.9999E-01 
 1.0000E-08 , 1.1510E-04 
 1.7500E-08 , 2.0622E-04 
`,
			`sweep 4
 2.5000E-09 , 8.0000E-01 
 1.0000E-08 , 1.0132E-04 
 1.7500E-08 , 2.0619E-04 
`,
			`sweep 5
 2.5000E-09 , 7.9998E-01 
 1.0000E-08 , 2.2212E-05 
 1.7500E-08 , 3.0525E-05 
`,
			`sweep 6
 2.5000E-09 , 8.0000E-01 
 1.0000E-08 , 5.9489E-04 
 1.7500E-08 , 1.1026E-03 
`,
			`sweep 7
 2.5000E-09 , 8.0002E-01 
 1.0000E-08 , 1.1060E-04 
 1.7500E-08 , 1.0251E-04 
`,
		},
	}

	assert.Equal(t, expect, actual)
}

func TestNewHeader(t *testing.T) {
	expect := Header{
		Format:    "2dsweep MONTE_CARLO",
		Signal:    "m8d",
		SavedTime: "15:12:20 Mon Feb 25 2019",
	}
	actual := NewHeader(`#format 2dsweep MONTE_CARLO
#[Custom WaveView] saved 15:12:20 Mon Feb 25 2019
TIME ,m8d 
`)
	assert.Equal(t, expect, actual)

}

func TestWVParser_Parse(t *testing.T) {
	actual, err := WVParser{FilePath: filepath.Join("test", "target.csv")}.Parse()
	assert.NoError(t, err)
	expect := WVCsv{
		Header: Header{
			Format:    "2dsweep MONTE_CARLO",
			Signal:    "m8d",
			SavedTime: "15:12:20 Mon Feb 25 2019",
		},
		Data: []*Element{
			{
				Name: "sweep 1",
				Values: map[float64]float64{2.5000E-09: 8.0001E-01,
					1.0000E-08: 1.0535E-04,
					1.7500E-08: 1.6397E-04,
				}},
			{
				Name: "sweep 2",
				Values: map[float64]float64{2.5000E-09: 8.0000E-01,
					1.0000E-08: 1.9597E-04,
					1.7500E-08: 2.6414E-04,
				}}, {Name: "sweep 3",
				Values: map[float64]float64{2.5000E-09: 7.9999E-01,
					1.0000E-08: 1.1510E-04,
					1.7500E-08: 2.0622E-04,
				}}, {Name: "sweep 4",
				Values: map[float64]float64{2.5000E-09: 8.0000E-01,
					1.0000E-08: 1.0132E-04,
					1.7500E-08: 2.0619E-04,
				}}, {Name: "sweep 5",
				Values: map[float64]float64{2.5000E-09: 7.9998E-01,
					1.0000E-08: 2.2212E-05,
					1.7500E-08: 3.0525E-05,
				}}, {Name: "sweep 6",
				Values: map[float64]float64{2.5000E-09: 8.0000E-01,
					1.0000E-08: 5.9489E-04,
					1.7500E-08: 1.1026E-03,
				}}, {Name: "sweep 7",
				Values: map[float64]float64{2.5000E-09: 8.0002E-01,
					1.0000E-08: 1.1060E-04,
					1.7500E-08: 1.0251E-04,
				},
			},
		},
	}

	assert.Equal(t, expect, actual)
}

func TestDocument_GetElements(t *testing.T) {
	expect := []*Element{
		{
			Name: "sweep 1",
			Values: map[float64]float64{2.5000E-09: 8.0001E-01,
				1.0000E-08: 1.0535E-04,
				1.7500E-08: 1.6397E-04,
			}},
		{
			Name: "sweep 2",
			Values: map[float64]float64{2.5000E-09: 8.0000E-01,
				1.0000E-08: 1.9597E-04,
				1.7500E-08: 2.6414E-04,
			}}, {Name: "sweep 3",
			Values: map[float64]float64{2.5000E-09: 7.9999E-01,
				1.0000E-08: 1.1510E-04,
				1.7500E-08: 2.0622E-04,
			}}, {Name: "sweep 4",
			Values: map[float64]float64{2.5000E-09: 8.0000E-01,
				1.0000E-08: 1.0132E-04,
				1.7500E-08: 2.0619E-04,
			}}, {Name: "sweep 5",
			Values: map[float64]float64{2.5000E-09: 7.9998E-01,
				1.0000E-08: 2.2212E-05,
				1.7500E-08: 3.0525E-05,
			}}, {Name: "sweep 6",
			Values: map[float64]float64{2.5000E-09: 8.0000E-01,
				1.0000E-08: 5.9489E-04,
				1.7500E-08: 1.1026E-03,
			}}, {Name: "sweep 7",
			Values: map[float64]float64{2.5000E-09: 8.0002E-01,
				1.0000E-08: 1.1060E-04,
				1.7500E-08: 1.0251E-04,
			},
		},
	}
	d,err := WVParser{FilePath:filepath.Join("test","target.csv")}.GetDocument()
	assert.NoError(t,err)
	actual,err := d.GetElements()
	assert.NoError(t, err)

	assert.Equal(t,expect,actual)


}
