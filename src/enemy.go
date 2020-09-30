
package main

type enemy struct {
	name string
	atk int
	hp  int
	lev int
	exp int
}

var enemy_cont []enemy =
	[]enemy{
		enemy{"Vi"		,	8,    30,    1,    2},
		enemy{"Git"		,	12,   20,    1,    3},
		enemy{"Python"	,	18,   40,    2,    5},
		enemy{"Vim"		,   15,   50,    3,    6},
		enemy{"Xfce"	,   20,   51,    4,    7},
		enemy{"Gnome"	,   23,   68,    6,   12},
		enemy{"Gentoo"	,   24,   90,    7,   17},
		enemy{"Xamarin"	,   31,   92,    8,   20},
		enemy{"Cygwin"	,   30,  100,    9,   25},
		enemy{"Chainer"	,   35,  125,   10,   26},
		enemy{"Django"  ,   32,  170,   11,   30},
		enemy{"Vulkan"	,   39,  148,   12,   32},
		enemy{"Firefox"	,   53,  110,   13,   40},
		enemy{"Doxygen"	,   43,  161,   14,   43},
		enemy{"Debian"	,   48,  170,   16,   78},
		enemy{"Fortran"	,   66,  150,   17,   80},
		enemy{"Julia"	,   60,  200,   18,  100},
		enemy{"Tux"     ,   65,  500,   40,  200},
		enemy{"Linus"   ,   70,  640,   50,    0},
	}