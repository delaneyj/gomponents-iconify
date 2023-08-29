package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelScheduleSendOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.975 17.725L18.35 19.1q.15.15.375.15t.375-.15q.15-.15.15-.375t-.15-.375L17.75 17l1.375-1.375q.15-.15.15-.375t-.15-.375q-.15-.15-.375-.15t-.375.15L17 16.25l-1.375-1.375q-.15-.15-.375-.15t-.375.15q-.15.15-.15.375t.15.375L16.25 17l-1.4 1.4q-.15.15-.15.35t.15.35q.15.15.375.15t.375-.15l1.375-1.375ZM3 18.5v-13q0-.55.438-.838t.937-.087L17.3 10H17q-.875 0-1.65.2t-1.5.55L5 7v3.5l6 1.5l-6 1.5V17l5.4-2.3q-.2.575-.3 1.137T10 17v.05l-5.6 2.375q-.5.2-.95-.088T3 18.5ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22ZM5 14.7V7v10v-2.3Z"/>`),
		g.Group(children),
	)
}