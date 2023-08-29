package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifeFitsIntoHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.05 20.04L25.632 6.191a2.531 2.531 0 0 0-3.263 0L5.954 20.04a1.266 1.266 0 0 0 .816 2.233h4.343v17.45a1 1 0 0 0 1 1h8.296v-6.209a3.73 3.73 0 0 1 3.4-3.792a3.594 3.594 0 0 1 3.786 3.589v6.412h8.296a1 1 0 0 0 1-1v-17.45h4.342a1.266 1.266 0 0 0 .816-2.233Z"/><circle cx="24.002" cy="23.159" r="4.506" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}