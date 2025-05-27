package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Theme struct {
	rosewater tcell.Color
	flamingo  tcell.Color
	pink      tcell.Color
	mauve     tcell.Color
	red       tcell.Color
	maroon    tcell.Color
	peach     tcell.Color
	yellow    tcell.Color
	green     tcell.Color
	teal      tcell.Color
	sky       tcell.Color
	sapphire  tcell.Color
	blue      tcell.Color
	lavender  tcell.Color
	text      tcell.Color
	subtext1  tcell.Color
	subtext0  tcell.Color
	overlay2  tcell.Color
	overlay1  tcell.Color
	overlay0  tcell.Color
	surface2  tcell.Color
	surface1  tcell.Color
	surface0  tcell.Color
	base      tcell.Color
	mantle    tcell.Color
	crust     tcell.Color
}

func MochaTheme() Theme {
	theme := Theme{
		rosewater: rgb(245, 224, 220),
		flamingo:  rgb(242, 205, 205),
		pink:      rgb(245, 194, 231),
		mauve:     rgb(203, 166, 247),
		red:       rgb(243, 139, 168),
		maroon:    rgb(235, 160, 172),
		peach:     rgb(250, 179, 135),
		yellow:    rgb(249, 226, 175),
		green:     rgb(166, 227, 161),
		teal:      rgb(148, 226, 213),
		sky:       rgb(137, 220, 235),
		sapphire:  rgb(116, 199, 236),
		blue:      rgb(137, 180, 250),
		lavender:  rgb(180, 190, 254),
		text:      rgb(205, 214, 244),
		subtext1:  rgb(186, 194, 222),
		subtext0:  rgb(166, 173, 200),
		overlay2:  rgb(147, 153, 178),
		overlay1:  rgb(127, 132, 156),
		overlay0:  rgb(108, 112, 134),
		surface2:  rgb(88, 91, 112),
		surface1:  rgb(69, 71, 90),
		surface0:  rgb(49, 50, 68),
		base:      rgb(30, 30, 46),
		mantle:    rgb(24, 24, 37),
		crust:     rgb(17, 17, 27),
	}
	updateGlobals(theme)
	return theme
}

func (theme Theme) Base() tcell.Color {
	return theme.base
}
func (theme Theme) Rosewater() tcell.Color {
	return theme.rosewater
}

func updateGlobals(theme Theme) {
	tview.Styles.PrimitiveBackgroundColor = theme.crust
	//tview.Styles.ContrastBackgroundColor = theme.base
	//tview.Styles.MoreContrastBackgroundColor = theme.base
	//tview.Styles.BorderColor = theme.base
	//tview.Styles.TitleColor = theme.base
	//tview.Styles.GraphicsColor = theme.base
	//tview.Styles.PrimaryTextColor = theme.base
	//tview.Styles.SecondaryTextColor = theme.base
	//tview.Styles.TertiaryTextColor = theme.base
	//tview.Styles.InverseTextColor = theme.base
	//tview.Styles.ContrastSecondaryTextColor = theme.base
}

func rgb(r, g, b int32) tcell.Color {
	return tcell.NewRGBColor(r, g, b)
}
