package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Healthelife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="20.51" cy="23.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.22" ry="10.21" transform="rotate(-20 20.523 23.56)"/><ellipse cx="14.75" cy="28.84" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.22" ry="10.21" transform="rotate(-65 14.751 28.839)"/><ellipse cx="27.49" cy="23.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.21" ry="4.22" transform="rotate(-70 27.495 23.562)"/><ellipse cx="33.25" cy="28.84" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.21" ry="4.22" transform="rotate(-25 33.242 28.831)"/>`),
		g.Group(children),
	)
}