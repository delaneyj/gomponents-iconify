package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path fill="#43CCF8" stroke="#fff" d="M29 14V34C25.15 34 22.2892 28.9106 22.2892 28.9106H18.1C17.4925 28.9106 17 28.4132 17 27.7995V20.1171C17 19.5034 17.4925 19.006 18.1 19.006H22.2892C22.2892 19.006 25.15 14 29 14Z"/></g>`),
		g.Group(children),
	)
}