package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cemetery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.46 12h-.68L12 3.55a.52.52 0 0 0-.56-.55h-1.18c0-.92-1.23-2-2.75-2S4.77 2.08 4.77 3H3.54a.52.52 0 0 0-.54.55L4.2 12h-.65a.53.53 0 0 0-.55.5V14h9v-1.51a.52.52 0 0 0-.54-.49zM4.5 5h6v1h-6V5z"/>`),
		g.Group(children),
	)
}