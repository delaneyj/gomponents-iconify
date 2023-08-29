package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M9 255L200 17c17-22 45-22 62 0l191 238c17 22 9 40-20 40H29c-29 0-37-18-20-40zm444 188L262 681c-17 22-45 22-62 0L9 443c-17-22-9-40 20-40h404c29 0 37 18 20 40z"/>`),
		g.Group(children),
	)
}