package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.8 2.515a2 2 0 0 1 2.494.476L21 8.638V22H3V8.406l10.8-5.89ZM5 10v4h14v-4H5Zm12.865-2l-3.107-3.73L7.922 8h9.943ZM19 16H5v4h14v-4Z"/>`),
		g.Group(children),
	)
}