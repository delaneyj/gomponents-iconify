package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 204h-12V96a20 20 0 0 0-20-20h-36V32a20 20 0 0 0-28.45-18.12l-104 48.54A20.06 20.06 0 0 0 28 80.55V204H16a12 12 0 0 0 0 24h224a12 12 0 0 0 0-24Zm-36-104v104h-32V100ZM52 83.09l96-44.79V204H52ZM132 112v12a12 12 0 0 1-24 0v-12a12 12 0 0 1 24 0Zm-40 0v12a12 12 0 0 1-24 0v-12a12 12 0 0 1 24 0Zm0 52v12a12 12 0 0 1-24 0v-12a12 12 0 0 1 24 0Zm40 0v12a12 12 0 0 1-24 0v-12a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}