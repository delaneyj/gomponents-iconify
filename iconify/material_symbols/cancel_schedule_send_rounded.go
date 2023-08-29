package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelScheduleSendRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.975 17.725L18.35 19.1q.15.15.375.15t.375-.15q.15-.15.15-.375t-.15-.375L17.75 17l1.375-1.375q.15-.15.15-.375t-.15-.375q-.15-.15-.375-.15t-.375.15L17 16.25l-1.375-1.375q-.15-.15-.375-.15t-.375.15q-.15.15-.15.375t.15.375L16.25 17l-1.4 1.4q-.15.15-.15.35t.15.35q.15.15.375.15t.375-.15l1.375-1.375ZM3 18.5V14l8-2l-8-2V5.5q0-.55.438-.838t.937-.087L17.3 10H17q-2.925 0-4.963 2.063T10 17.05v-.025v.025l-5.6 2.375q-.5.2-.95-.088T3 18.5ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Z"/>`),
		g.Group(children),
	)
}