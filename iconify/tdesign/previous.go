package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.25 4.336v15.328L9.586 12l7.664-7.664ZM8.5 5v14h-2V5h2Zm3.914 7l2.836 2.836V9.164L12.414 12Z"/>`),
		g.Group(children),
	)
}