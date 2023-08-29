package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 10a3 3 0 0 0-3 3v4h40v-4a3 3 0 0 0-3-3H7Zm37 13H4v12a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V23Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}