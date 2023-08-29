package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneHexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.05 19h-8.1l-4.04-7l4.04-7h8.1l4.04 7z" opacity=".3"/><path fill="currentColor" d="M17.2 3H6.8l-5.2 9l5.2 9h10.4l5.2-9l-5.2-9zm-1.15 16h-8.1l-4.04-7l4.04-7h8.09l4.04 7l-4.03 7z"/>`),
		g.Group(children),
	)
}