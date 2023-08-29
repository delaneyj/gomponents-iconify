package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H48a14 14 0 0 0-14 14v58.77c0 88.24 74.68 117.52 89.65 122.49a13.5 13.5 0 0 0 8.7 0c15-5 89.65-34.25 89.65-122.49V56a14 14 0 0 0-14-14Zm2 72.79c0 80-67.84 106.59-81.44 111.1a1.55 1.55 0 0 1-1.12 0c-13.6-4.51-81.44-31.1-81.44-111.1V56a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}