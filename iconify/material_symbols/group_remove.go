package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 11.95q.725-.8 1.113-1.825T14 8q0-1.1-.388-2.125T12.5 4.05q1.5.2 2.5 1.325T16 8q0 1.5-1 2.625t-2.5 1.325ZM18 20v-3q0-.9-.4-1.713t-1.05-1.437q1.275.45 2.363 1.163T20 17v3h-2Zm6-9h-6V9h6v2ZM8 12q-1.65 0-2.825-1.175T4 8q0-1.65 1.175-2.825T8 4q1.65 0 2.825 1.175T12 8q0 1.65-1.175 2.825T8 12Zm-8 8v-2.8q0-.85.438-1.563T1.6 14.55q1.55-.775 3.15-1.163T8 13q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T16 17.2V20H0Z"/>`),
		g.Group(children),
	)
}