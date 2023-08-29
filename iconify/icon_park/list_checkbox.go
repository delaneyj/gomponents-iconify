package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" fill-rule="evenodd" d="M20 24H44H20Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 24H44"/><path fill="#2F88FF" fill-rule="evenodd" d="M20 38H44H20Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 38H44"/><path fill="#2F88FF" fill-rule="evenodd" d="M20 10H44H20Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 10H44"/><rect width="8" height="8" x="4" y="34" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/><rect width="8" height="8" x="4" y="20" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/><rect width="8" height="8" x="4" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}