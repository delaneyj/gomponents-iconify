package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cncf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h24V0H8.004Zm3.431 3.431h4.544l.029.023l4.544 4.544h3.03l-4.572-4.567h9.569v9.563l-.789-.782l-3.784-3.79v3.03l2.271 2.272l2.272 2.272l.029.03v4.543h-4.55l-.023-.023l-2.272-2.278l-2.272-2.272h-3.03l3.785 3.79l.782.783H3.43v-9.563l4.573 4.567v-3.031l-4.55-4.544l-.023-.023Z"/>`),
		g.Group(children),
	)
}