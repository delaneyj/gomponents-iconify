package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2H3v20h18V2H4zm15 2v16H5V4h14zm-6 2h-2v2h2V6zm-5 4h8v6h-2v-4h-4v4H8v-6zm8 6H8v2h8v-2z"/>`),
		g.Group(children),
	)
}