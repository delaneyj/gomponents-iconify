package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M10 4H30L40 14V42C40 43.1046 39.1046 44 38 44H10C8.89543 44 8 43.1046 8 42V6C8 4.89543 8.89543 4 10 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M18 18H30V25.9917L18.0083 26L18 18Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M18 18V34"/></g>`),
		g.Group(children),
	)
}