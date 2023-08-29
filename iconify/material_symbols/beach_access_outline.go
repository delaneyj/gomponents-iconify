package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeachAccessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.6 21l-6.35-6.35l1.4-1.4L21 19.6L19.6 21Zm-13.65-.7q-1.5-1.5-2.225-3.375T3 13.1q0-1.95.725-3.8T5.95 5.95q1.5-1.5 3.363-2.237t3.812-.738q1.95 0 3.813.738T20.3 5.95L5.95 20.3Zm.2-3.05L7.5 15.9q-.4-.525-.762-1.075t-.663-1.1q-.3-.55-.525-1.1t-.4-1.075q-.275 1.475-.038 2.95t1.038 2.75Zm2.8-2.75l5.55-5.6q-1.075-.825-2.162-1.338t-2.038-.7q-.95-.187-1.712-.062T7.4 7.35q-.425.45-.55 1.213t.063 1.724q.187.963.7 2.038T8.95 14.5Zm6.95-7l1.4-1.35q-1.325-.8-2.8-1.05t-2.95.05q.55.175 1.1.4t1.1.513q.55.287 1.088.65T15.9 7.5Z"/>`),
		g.Group(children),
	)
}