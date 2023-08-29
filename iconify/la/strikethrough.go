package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 6v2h8v8h-5v2h5v8h2v-8h5v-2h-5V8h8V6z"/>`),
		g.Group(children),
	)
}