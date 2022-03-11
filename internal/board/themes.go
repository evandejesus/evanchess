package board

import "image/color"

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
