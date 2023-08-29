package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDolbyDigital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6v12h-.89c-3.34 0-6.047-2.686-6.047-6s2.707-6 6.046-6H21zM3.063 6v12h.891C7.294 18 10 15.314 10 12S7.293 6 3.954 6h-.89z"/>`),
		g.Group(children),
	)
}