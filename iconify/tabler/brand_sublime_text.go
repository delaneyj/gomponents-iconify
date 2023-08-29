package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSublimeText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 8L5 12.5V7l14-4.5zm0 9L5 21.5V16l14-4.5zm0-5.5L5 7m0 5.5L19 17"/>`),
		g.Group(children),
	)
}