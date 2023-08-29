package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 50H32a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h192a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14Zm2 142a2 2 0 0 1-2 2H32a2 2 0 0 1-2-2V64a2 2 0 0 1 2-2h192a2 2 0 0 1 2 2ZM50 136a6 6 0 0 1 6-6h16a6 6 0 0 1 0 12H56a6 6 0 0 1-6-6Zm156 0a6 6 0 0 1-6 6h-96a6 6 0 0 1 0-12h96a6 6 0 0 1 6 6Zm-48 32a6 6 0 0 1-6 6H56a6 6 0 0 1 0-12h96a6 6 0 0 1 6 6Zm48 0a6 6 0 0 1-6 6h-16a6 6 0 0 1 0-12h16a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}