package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFlipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.973 3h16.054c.537 0 .973.436.973.973v4.052a.973.973 0 0 1-.973.973h-5.025v4.831c0 .648-.525 1.173-1.173 1.173H9v5.025a.973.973 0 0 1-.974.973H3.973A.973.973 0 0 1 3 20.027V3.973C3 3.436 3.436 3 3.973 3z"/>`),
		g.Group(children),
	)
}