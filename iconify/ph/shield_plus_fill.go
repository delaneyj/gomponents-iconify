package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldPlusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 40H48a16 16 0 0 0-16 16v58.77c0 89.62 75.82 119.34 91 124.38a15.44 15.44 0 0 0 10 0c15.2-5.05 91-34.77 91-124.39V56a16 16 0 0 0-16-16Zm-48 96h-24v24a8 8 0 0 1-16 0v-24H96a8 8 0 0 1 0-16h24V96a8 8 0 0 1 16 0v24h24a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}