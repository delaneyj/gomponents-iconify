package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScheduleSendRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19q.15-.15.15-.35T19 18.3l-1.5-1.5v-2.3q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35v2.275q0 .2.075.387t.225.338l1.5 1.5q.15.15.35.15T19 19Zm-16-.5V14l8-2l-8-2V5.5q0-.55.438-.838t.937-.087L17.3 10H17q-2.925 0-4.963 2.063T10 17.05v-.025v.025l-5.6 2.375q-.5.2-.95-.088T3 18.5ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Z"/>`),
		g.Group(children),
	)
}