package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 160h-8v-39.5c1.63.81 3.29 1.57 5 2.26a8 8 0 0 0 6-14.83A55.78 55.78 0 0 1 200 56a8 8 0 0 0-16 0a56 56 0 0 1-112 0a8 8 0 0 0-16 0a55.78 55.78 0 0 1-35 51.93a8 8 0 0 0 6 14.83c1.71-.69 3.37-1.45 5-2.26V160h-7.4c-6.31 0-8.6 4.78-8.6 8a8 8 0 0 0 8 8h32v24a8 8 0 0 0 16 0v-24h112v24a8 8 0 0 0 16 0v-24h32a8 8 0 0 0 0-16Zm-160-8a8 8 0 0 1-16 0v-44.65a8 8 0 0 1 16 0Zm40 0a8 8 0 0 1-16 0v-20.65a8 8 0 0 1 16 0Zm48 0a8 8 0 0 1-16 0v-20.65a8 8 0 0 1 16 0Zm40 0a8 8 0 0 1-16 0v-44.65a8 8 0 0 1 16 0Zm0-50.2v-.46l.41.51Z"/>`),
		g.Group(children),
	)
}