package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func R(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-5.5-7.362l3.467-1.812V10.745l4.952 2.778l-3.714 1.933v3.987L23.5 25v-3.745l-5.076-3.503l4.21-2.175v-3.866L13.966 7L10.5 8.812z"/>`),
		g.Group(children),
	)
}