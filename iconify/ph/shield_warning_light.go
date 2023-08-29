package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldWarningLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H48a14 14 0 0 0-14 14v58.77c0 88.25 74.68 117.52 89.65 122.49a13.5 13.5 0 0 0 8.7 0c15-5 89.65-34.24 89.65-122.49V56a14 14 0 0 0-14-14Zm2 72.79c0 80-67.84 106.59-81.44 111.1a1.57 1.57 0 0 1-1.13 0C113.84 221.38 46 194.8 46 114.79V56a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2ZM122 136V96a6 6 0 0 1 12 0v40a6 6 0 0 1-12 0Zm16 36a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}