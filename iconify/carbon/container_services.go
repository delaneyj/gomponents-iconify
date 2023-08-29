package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22v-5a2.002 2.002 0 0 0-2-2h-8v-5h3a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2h-8a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h3v5H7a2.002 2.002 0 0 0-2 2v5H2v8h8v-8H7v-5h8v5h-3v8h8v-8h-3v-5h8v5h-3v8h8v-8ZM12 4h8v4h-8ZM8 28H4v-4h4Zm10 0h-4v-4h4Zm10 0h-4v-4h4Z"/>`),
		g.Group(children),
	)
}