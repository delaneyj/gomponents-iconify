package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 44H44"/><rect width="12" height="12" x="27" y="32" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 10V16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 10V16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 10L33 4L38 10H28Z"/><path stroke="#000" stroke-linejoin="round" stroke-width="4" d="M25 20H11C9.89543 20 9 20.8954 9 22V44"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M15 29L19 29"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M15 36L19 36"/><rect width="16" height="16" x="25" y="16" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="33" cy="24" r="3" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M33 32V42"/></g>`),
		g.Group(children),
	)
}