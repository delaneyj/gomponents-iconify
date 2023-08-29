package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.22 9.5a6.219 6.219 0 0 0-5.885 4.218H1.438v4h1.897a6.216 6.216 0 0 0 12.102-2A6.217 6.217 0 0 0 9.218 9.5zm18.465 4.22c-.944-5.173-5.46-9.095-10.903-9.095v4a7.104 7.104 0 0 1 7.094 7.094a7.106 7.106 0 0 1-7.094 7.092v4.002c5.442-.004 9.96-3.926 10.903-9.096h2.065v-4h-2.065z"/>`),
		g.Group(children),
	)
}