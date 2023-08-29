package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendAndArchiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18.5V14l8-2l-8-2V5.5q0-.55.438-.838t.937-.087L17.3 10H17q-2.925 0-4.963 2.063T10 17.05v-.025v.025l-5.6 2.375q-.5.2-.95-.088T3 18.5ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm-.5-3.9l-1.45-1.45q-.15-.15-.35-.15t-.35.15q-.15.15-.15.35t.15.35l1.95 1.95q.3.3.7.3t.7-.3l1.95-1.95q.15-.15.15-.35t-.15-.35q-.15-.15-.35-.15t-.35.15L17.5 18.1v-3.6q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35v3.6Z"/>`),
		g.Group(children),
	)
}