package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50h-64a38 38 0 0 0-32 17.55A38 38 0 0 0 96 50H32a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h64a26 26 0 0 1 26 26a6 6 0 0 0 12 0a26 26 0 0 1 26-26h64a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14ZM96 194H32a2 2 0 0 1-2-2V64a2 2 0 0 1 2-2h64a26 26 0 0 1 26 26v116.31A37.86 37.86 0 0 0 96 194Zm130-2a2 2 0 0 1-2 2h-64a37.87 37.87 0 0 0-26 10.32V88a26 26 0 0 1 26-26h64a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}