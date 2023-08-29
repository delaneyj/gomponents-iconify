package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlackLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M218 128a30 30 0 1 0-36-48V56a30 30 0 0 0-54-18a30 30 0 1 0-48 36H56a30 30 0 0 0-18 54a30 30 0 1 0 36 48v24a30 30 0 0 0 54 18a30 30 0 1 0 48-36h24a30 30 0 0 0 18-54Zm-18-42a18 18 0 0 1 0 36h-18v-18a18 18 0 0 1 18-18Zm-48-48a18 18 0 0 1 18 18v48a18 18 0 0 1-18 18h-18V56a18 18 0 0 1 18-18ZM86 56a18 18 0 0 1 36 0v18h-18a18 18 0 0 1-18-18Zm-48 48a18 18 0 0 1 18-18h48a18 18 0 0 1 18 18v18H56a18 18 0 0 1-18-18Zm18 66a18 18 0 0 1 0-36h18v18a18 18 0 0 1-18 18Zm48 48a18 18 0 0 1-18-18v-48a18 18 0 0 1 18-18h18v66a18 18 0 0 1-18 18Zm66-18a18 18 0 0 1-36 0v-18h18a18 18 0 0 1 18 18Zm30-30h-48a18 18 0 0 1-18-18v-18h66a18 18 0 0 1 0 36Z"/>`),
		g.Group(children),
	)
}