package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemplateOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h11a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-7M8 8H5a1 1 0 0 1-1-1V5c0-.271.108-.517.283-.697M4 13a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1zm12-1h4m-6 4h2m-2 4h6M3 3l18 18"/>`),
		g.Group(children),
	)
}