package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2h-2v2H8v2H6v2H2v8h4v2h2v2h2v2h2V2zM8 18v-2H6v-2H4v-4h2V8h2V6h2v12H8zm14-7h-8v2h8v-2z"/>`),
		g.Group(children),
	)
}