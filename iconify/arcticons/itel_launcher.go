package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItelLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="22.248" cy="10.214" r="5.715" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.747 19.456h4.005c4.144 0 6.906 2.366 7.422 6.345L31.466 43.5H20.863l-3.116-24.044Z"/>`),
		g.Group(children),
	)
}