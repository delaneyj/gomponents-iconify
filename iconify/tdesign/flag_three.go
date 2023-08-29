package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 10L8 18.66V1.34L23 10ZM10 4.804v10.392L19 10l-9-5.196ZM4 1h2v22H4V1Z"/>`),
		g.Group(children),
	)
}