package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 1 1 1v1h1a1 1 0 0 1 0 2H6v1a1 1 0 0 1-2 0V6H3a1 1 0 0 1 0-2h1V3a1 1 0 0 1 1-1Zm0 10a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2H6v1a1 1 0 1 1-2 0v-1H3a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1Zm7-10a1 1 0 0 1 .967.744L14.146 7.2L17.5 9.134a1 1 0 0 1 0 1.732l-3.354 1.935l-1.18 4.455a1 1 0 0 1-1.933 0L9.854 12.8L6.5 10.866a1 1 0 0 1 0-1.732l3.354-1.935l1.18-4.455A1 1 0 0 1 12 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}