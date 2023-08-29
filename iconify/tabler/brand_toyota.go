package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandToyota(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a10 7 0 1 0 20 0a10 7 0 1 0-20 0"/><path d="M9 12c0 3.866 1.343 7 3 7s3-3.134 3-7s-1.343-7-3-7s-3 3.134-3 7z"/><path d="M6.415 6.191C5.527 6.694 5 7.321 5 8c0 1.657 3.134 3 7 3s7-1.343 7-3c0-.678-.525-1.304-1.41-1.806"/></g>`),
		g.Group(children),
	)
}