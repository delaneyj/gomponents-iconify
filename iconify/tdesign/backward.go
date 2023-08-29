package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 4.336v7l7-7v15.328l-7-7v7L3.586 12l7.664-7.664ZM6.414 12l2.836 2.836V9.164L6.414 12Zm7 0l2.836 2.836V9.164L13.414 12Z"/>`),
		g.Group(children),
	)
}