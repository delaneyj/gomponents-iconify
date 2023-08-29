package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .932l5.288 1.983l-.703 1.873l-.585-.22V5h3v12h-2v6H7v-6H5V5h3v-.432l-.585.22l-.702-1.873L12 .932Zm-2 2.886V5h4V3.818l-2-.75l-2 .75ZM9 17v4h6v-4h-2v3h-2v-3H9Zm8-2V7H7v8h10Zm-5-5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}