package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4h-2v16h2V4zM4 6h5v2H4v8h5v2H2V6h2zm11 0h7v12h-7v-2h5V8h-5V6z"/>`),
		g.Group(children),
	)
}