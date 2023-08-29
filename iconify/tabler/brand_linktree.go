package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandLinktree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3L2 15h3v5h5v-5H8l4-7zm6 0l7 12h-3v5h-5v-5h2l-4-7z"/>`),
		g.Group(children),
	)
}