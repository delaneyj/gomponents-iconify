package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2v2h5v2h-8V2h3zM5 16h9v2h2v4h-6v-2h4v-2h-4v4H8v-4H3v-6h2v-2h2v4H5v2zm16-2h-2v-2h-2V6h2v6h2v2zM5 2H3v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2V8H9V6H7V4H5V2z"/>`),
		g.Group(children),
	)
}