package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDocker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12.54c-1.804-.345-2.701-1.08-3.523-2.94c-.487.696-1.102 1.568-.92 2.4c.028.238-.32 1-.557 1H3c0 5.208 3.164 7 6.196 7c4.124.022 7.828-1.376 9.854-5c1.146-.101 2.296-1.505 2.95-2.46z"/><path d="M5 10h3v3H5zm3 0h3v3H8zm3 0h3v3h-3zM8 7h3v3H8zm3 0h3v3h-3zm0-3h3v3h-3zM4.571 18c1.5 0 2.047-.074 2.958-.78M10 16v.01"/></g>`),
		g.Group(children),
	)
}