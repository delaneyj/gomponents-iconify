package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCodecov(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.695 12.985A5.972 5.972 0 0 0 6.4 12c-1.257 0-2.436.339-3.4 1a9 9 0 1 1 18 0c-.966-.664-2.14-1-3.4-1a6 6 0 0 0-5.605 8.144"/>`),
		g.Group(children),
	)
}