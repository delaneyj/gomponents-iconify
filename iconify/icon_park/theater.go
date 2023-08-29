package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke-linejoin="round" d="M8 6H40C41.1046 6 42 6.89543 42 8V36L32 22.0053C29.3333 23.5901 26.6667 24.3825 24 24.3825C21.3333 24.3825 18.6667 23.5901 16 22.0053L6 36V8C6 6.89543 6.89543 6 8 6Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M6 42H42"/></g>`),
		g.Group(children),
	)
}