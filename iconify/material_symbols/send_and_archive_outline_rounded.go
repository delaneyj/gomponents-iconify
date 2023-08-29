package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendAndArchiveOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 18.1l-1.45-1.45q-.15-.15-.35-.15t-.35.15q-.15.15-.15.35t.15.35l1.95 1.95q.3.3.7.3t.7-.3l1.95-1.95q.15-.15.15-.35t-.15-.35q-.15-.15-.35-.15t-.35.15L17.5 18.1v-3.6q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35v3.6ZM3 18.5v-13q0-.55.438-.838t.937-.087L17.3 10H17q-.875 0-1.65.2t-1.5.55L5 7v3.5l6 1.5l-6 1.5V17l5.4-2.3q-.2.575-.3 1.137T10 17v.05l-5.6 2.375q-.5.2-.95-.088T3 18.5ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22ZM5 14.7V7v10v-2.3Z"/>`),
		g.Group(children),
	)
}