package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abacus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm14 2v3h-.1c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H7.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V2h12zm-.1 8c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H4.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V6h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1l-.1 4zM2 14v-3h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1v3H2z"/>`),
		g.Group(children),
	)
}