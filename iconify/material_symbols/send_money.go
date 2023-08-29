package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendMoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19.75q-2.6-.675-4.3-2.8T0 12q0-2.825 1.7-4.95T6 4.25v2.1q-1.775.6-2.888 2.15T2 12q0 1.95 1.113 3.5T6 17.65v2.1Zm8 .25q-3.325 0-5.663-2.337T6 12q0-3.325 2.337-5.663T14 4q1.65 0 3.1.625t2.55 1.725l-1.4 1.4q-.825-.825-1.912-1.288T14 6q-2.5 0-4.25 1.75T8 12q0 2.5 1.75 4.25T14 18q1.25 0 2.337-.463t1.913-1.287l1.4 1.4q-1.1 1.1-2.55 1.725T14 20Zm6-4l-1.4-1.4l1.6-1.6H13v-2h7.2l-1.6-1.6L20 8l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}