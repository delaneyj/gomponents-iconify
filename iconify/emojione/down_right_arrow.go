package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M21.9 44.5L44 44l.5-22.1l-7.5 7.5l-13.2-13.3l-7.3 7.3l13.3 13.2z"/>`),
		g.Group(children),
	)
}