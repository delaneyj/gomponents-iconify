package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.24 16v10.667L16 32V21.328l-9.24 5.339V5.328L16 0v10.667L6.76 16L16 21.328V10.667l9.24-5.339z"/>`),
		g.Group(children),
	)
}