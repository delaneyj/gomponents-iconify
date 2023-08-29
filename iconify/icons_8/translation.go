package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 4v18h6v6h18V10h-6V4H4zm2 2h14v4.563L10.562 20H6V6zm5 2v1H8v2h4.938c-.13 1.15-.482 2.054-1.063 2.688a4.567 4.567 0 0 1-.906-.407c-.704-.418-.97-.86-.97-1.28H8c0 1.192.734 2.182 1.72 2.844A8.487 8.487 0 0 1 8 15v2c1.772 0 3.248-.405 4.375-1.156c.524.09 1.053.156 1.625.156v-1.875c.543-.91.833-1.973.938-3.125H16V9h-3V8h-2zm10.438 4H26v14H12v-4.563L21.438 12zM20 13.844l-.938 2.844l-2 6l-.062.156V24h2v-.875l.03-.125h1.94l.03.125V24h2v-1.156l-.063-.157l-2-6L20 13.845zm0 6.28l.28.876h-.56l.28-.875z"/>`),
		g.Group(children),
	)
}