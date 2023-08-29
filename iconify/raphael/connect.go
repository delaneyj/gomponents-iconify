package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.06 13.72c-.944-5.173-5.46-9.095-10.903-9.095v4a7.104 7.104 0 0 1 7.094 7.094a7.104 7.104 0 0 1-7.093 7.092v4.002c5.442-.004 9.96-3.926 10.903-9.096h4.69v-4h-4.69zm-4.685 2a6.216 6.216 0 0 0-12.103-2.002H1.438v4h6.834a6.216 6.216 0 0 0 12.104-2z"/>`),
		g.Group(children),
	)
}