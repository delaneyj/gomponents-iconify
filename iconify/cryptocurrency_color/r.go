package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func R(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#771a4e"/><path fill="#fff" d="m10.5 24.638l3.467-1.812V10.745l4.952 2.778l-3.714 1.933v3.987L23.5 25v-3.745l-5.076-3.503l4.209-2.175v-3.866L13.967 7L10.5 8.812z"/></g>`),
		g.Group(children),
	)
}