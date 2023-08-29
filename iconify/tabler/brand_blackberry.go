package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBlackberry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6a1 1 0 0 0-1-1H4l-.5 2H6a1 1 0 0 0 1-1zm-1 6a1 1 0 0 0-1-1H3l-.5 2H5a1 1 0 0 0 1-1zm7 0a1 1 0 0 0-1-1h-2l-.5 2H12a1 1 0 0 0 1-1zm1-6a1 1 0 0 0-1-1h-2l-.5 2H13a1 1 0 0 0 1-1zm-2 12a1 1 0 0 0-1-1H9l-.5 2H11a1 1 0 0 0 1-1zm8-3a1 1 0 0 0-1-1h-2l-.5 2H19a1 1 0 0 0 1-1zm1-6a1 1 0 0 0-1-1h-2l-.5 2H20a1 1 0 0 0 1-1z"/>`),
		g.Group(children),
	)
}