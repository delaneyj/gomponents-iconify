package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeVibrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2h-2v2h-2v2H8v2H4v8h4v2h2v2h2v2h2V2zm-4 16v-2H8v-2H6v-4h2V8h2V6h2v12h-2zm8-15h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v-2h2v-2h-2v-2h2v-2h-2v-2h2V9h-2V7h2V5h-2V3z"/>`),
		g.Group(children),
	)
}