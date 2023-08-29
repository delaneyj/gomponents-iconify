package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuBurger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.563 4.063a.5.5 0 0 1 0-1l16.874-.001a.5.5 0 0 1 0 1l-16.874.001Zm0 8.438a.5.5 0 0 1 0-1l16.874-.002a.5.5 0 0 1 0 1l-16.874.002Zm0 8.438a.5.5 0 0 1 0-1l16.874-.002a.5.5 0 0 1 0 1l-16.874.002Z"/>`),
		g.Group(children),
	)
}