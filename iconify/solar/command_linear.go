package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M8 8h8v8H8V8Zm8 8.001h3a3 3 0 1 1-3 3v-3Zm-7.999 0h-3a3 3 0 1 0 3 3v-3ZM16 8h3a3 3 0 1 0-3-3v3ZM8.001 8h-3a3 3 0 1 1 3-3v3Z"/>`),
		g.Group(children),
	)
}