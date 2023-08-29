package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceTablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2H4v20h16V2H6zm12 2v16H6V4h12zm-5 12h-2v2h2v-2z"/>`),
		g.Group(children),
	)
}