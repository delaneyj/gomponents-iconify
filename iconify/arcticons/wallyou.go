package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallyou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.344 38.303H4.5l13.406-19.292l8.264 11.892"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.344 38.303l11.578-14.668L43.5 38.302H20.344ZM37.365 20.07l-1.47-1.618l-2.14-.04l.105-2.183l-1.485-1.542l1.618-1.47l.04-2.14l2.183.105l1.542-1.485l1.469 1.618l2.14.04l-.105 2.183l1.486 1.541l-1.618 1.47l-.04 2.14l-2.183-.105l-1.542 1.486Z"/><circle cx="37.589" cy="14.964" r="1.706" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}