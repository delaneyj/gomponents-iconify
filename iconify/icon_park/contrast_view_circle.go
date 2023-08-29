package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastViewCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44V4Z" clip-rule="evenodd"/><path fill="#2F88FF" d="M24 4C35.0457 4 44 12.9543 44 24C44 35.0457 35.0457 44 24 44V4Z"/><path stroke-linecap="round" d="M24 36H9"/><path stroke-linecap="round" d="M24 28H5"/><path stroke-linecap="round" d="M24 20H5"/><path stroke-linecap="round" d="M24 12H9"/></g>`),
		g.Group(children),
	)
}