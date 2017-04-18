package mahjan

import (
"fmt"
)

type Yaku struct {
	Name string
	Han uint
	Menzen bool
	Kuisagari bool
}

type Mahjan struct {
	Yakulist []Yaku
}

type Person uint

const (
	Parent Person = 0
	Child Person = 1
	)

func (this Yaku) String() string {
	kuisagari := func(y Yaku) string {
		if y.Kuisagari {
			return "/食い下がり"
		}
		return ""
	}(this)

	menzen := func(y Yaku) string {
		if y.Menzen {
			return "/門前のみ"
		}
		return ""
	}(this)

	han := fmt.Sprintf("%d翻",this.Han)
	switch this.Han {
		case 13: han = "役満"
		case 26: han = "ダブル役満"
	}

	return fmt.Sprintf("%s (%s%s%s)", this.Name, han, menzen, kuisagari)
}

func New() *Mahjan {
	yakulist := []Yaku {
		{"リーチ", 1, true, false},
		{"一発", 1, true, false},
		{"門前清自摸和", 1, true, false},
		{"平和", 1, true, false},
		{"断么九", 1, true, false},
		{"一盃口", 1, true, false},
		{"三元牌", 1, false, false},
		{"門風牌", 1, false, false},
		{"荘風牌", 1, false, false},
		{"嶺上開花", 1, false, false},
		{"槍槓", 1, false, false},
		{"海底撈月", 1, false, false},
		{"河底撈月", 1, false, false},
		{"ダブルリーチ", 2, true, false},
		{"全帯", 2, false, true},
		{"混老頭", 2, false, false},
		{"三色同順", 2, false, true},
		{"一気通貫", 2, false, true},
		{"対々和", 2, false, false},
		{"三色同刻", 2, false, false},
		{"三暗刻", 2, false, false},
		{"三槓子", 2, false, false},
		{"小三元", 2, false, false},
		{"七対子", 2, true, false},
		{"二盃口", 3, true, false},
		{"純全帯", 3, false, true},
		{"混一色", 3, false, true},
		{"清一色", 6, false, true},
		{"四暗刻", 13, true, false},
		{"四暗刻単騎待ち", 26, true, false},
		{"大三元", 13, false, false},
		{"字一色", 13, false, false},
		{"小四喜", 13, false, false},
		{"大四喜", 26, false, false},
		{"緑一色", 13, false, false},
		{"九蓮宝燈", 13, true, false},
		{"純正九蓮宝燈", 26, true, false},
		{"清老頭", 13, false, false},
		{"四槓子", 13, false, false},
		{"国士無双", 13, true, false},
		{"国士無双十三面待ち", 26, true, false},
		{"天和", 13, true, false},
		{"地和", 13, true, false},
	}
	
	return &Mahjan {
		Yakulist : yakulist,
	}
}

func (this *Mahjan) GetYakulist(han uint) []Yaku {
	yakulist := make([]Yaku, 0)

	for _, v := range this.Yakulist {
		if v.Han == han {
			yakulist = append(yakulist, v)
		}
	}

	return yakulist
}

func kiriage(score uint) uint {
	hasuu := score % 100
	if hasuu > 0 {
		return score + 100 - hasuu
	}
	return score
}

func (this *Mahjan) Score(hu, han uint, p Person, tsumo bool) string{

	if (hu % 10 != 0 && hu != 25) || (hu > 110 || hu < 20){
		return fmt.Sprintf("%d符は存在しないぞ", hu)
	}

	var tmp uint = 4
	if p == Parent {
		tmp = 6
	}

	score := hu * tmp
	switch {
		case (hu < 60 && han < 4) || (hu < 30 && han < 5):
			var loop int = int(han) + 2
			for ; loop > 0; loop-- {
				score *= 2
			}
		case han < 6:
			score = 2000 * tmp
		case han < 8:
			score = 3000 * tmp
		case han < 11:
			score = 4000 * tmp
		case han < 13:
			score = 6000 * tmp
		default:
			score = 8000 * tmp
	}

	if tsumo {
		switch p {
			case Parent :
				cscore := kiriage(score / 3)
				return fmt.Sprintf("%d All", cscore)
			case Child :
				pscore := kiriage(score / 2)
				cscore := kiriage(score / 4)
				return fmt.Sprintf("%d/%d", cscore, pscore)
		}
	}

	return fmt.Sprintf("%d", kiriage(score))
}
