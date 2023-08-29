package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16.2V7.8l6 4.2l-6 4.2Zm8-.2V8h2v8h-2Zm4 0V8h2v8h-2Z"/>`),
		g.Group(children),
	)
}