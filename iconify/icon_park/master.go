package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Master(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 38V6C38 4.89543 37.1046 4 36 4H12C10.8954 4 10 4.89543 10 6V38"/><rect width="28" height="12" x="10" y="32" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="6"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M20 16L24 12L28 16L32 12L29 23H19L16 12L20 16Z"/><circle cx="32" cy="38" r="2" fill="#fff"/></g>`),
		g.Group(children),
	)
}