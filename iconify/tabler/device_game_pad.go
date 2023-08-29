package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceGamePad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12L7 9H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2l3-3zm4 0l3-3h2a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2l-3-3zm-2 2l-3 3v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-2l-3-3zm0-4L9 7V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2l-3 3z"/>`),
		g.Group(children),
	)
}