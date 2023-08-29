package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h4v2H4v8h4v2H2V4h2zm6 0h10v2h-8v12h8v2H10V4zm12 0h-2v16h2V4zm-7 4h2v2h-2V8zm3 4h-4v4h4v-4zM8 18H4v2h4v-2z"/>`),
		g.Group(children),
	)
}