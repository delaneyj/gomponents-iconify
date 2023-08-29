package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Left(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M181 384q7 0 15-4q17-17 2-30L60 192L196 36q15-13-2-30q-14-14-30 3L4 192l162 186q4 6 15 6z"/>`),
		g.Group(children),
	)
}