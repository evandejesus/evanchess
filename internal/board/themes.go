/*
source: http://omgchess.blogspot.com/2015/09/chess-board-color-schemes.html
*/

package board

import "image/color"

//  Theme is a chess board color scheme giving colors for each player.
type Theme struct {
	light, dark color.RGBA
}

var Sandcastle = Theme{
	light: color.RGBA{227, 193, 111, 255},
	dark:  color.RGBA{184, 139, 74, 255},
}

var Emerald = Theme{
	light: color.RGBA{173, 189, 143, 255},
	dark:  color.RGBA{111, 143, 114, 255},
}

var Dusk = Theme{
	light: color.RGBA{204, 183, 174, 255},
	dark:  color.RGBA{112, 102, 119, 255},
}

var Coral = Theme{
	light: color.RGBA{177, 228, 185, 255},
	dark:  color.RGBA{112, 162, 163, 255},
}

var Wheat = Theme{
	light: color.RGBA{234, 240, 206, 255},
	dark:  color.RGBA{187, 190, 100, 255},
}

var Standard = Theme{
	light: color.RGBA{232, 231, 211, 255},
	dark:  color.RGBA{84, 113, 149, 255},
}

var Classic = Theme{
	light: color.RGBA{231, 202, 167, 255},
	dark:  color.RGBA{204, 165, 114, 255},
}

var Tournament = Theme{
	light: color.RGBA{231, 231, 231, 255},
	dark:  color.RGBA{62, 101, 76, 255},
}
