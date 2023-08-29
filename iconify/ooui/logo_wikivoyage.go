package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWikivoyage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.874 8.19L20 0l-8.254 4.095c.99.263 1.892.78 2.617 1.498a5.78 5.78 0 0 1 1.51 2.596Zm-6.349 4.043L14.678 20l.56-9.291a5.916 5.916 0 0 1-2.65 1.523a5.95 5.95 0 0 1-3.063 0Zm-.159-7.508L.001 5.28l7.829 5.113a5.813 5.813 0 0 1 0-3.038c.27-.995.8-1.902 1.536-2.63Z"/>`),
		g.Group(children),
	)
}