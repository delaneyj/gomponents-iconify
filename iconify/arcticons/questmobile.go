package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Questmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.033 8.667L4.5 39.333l39-.067Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.536 15.895c-5.95.628-15.44 15.854-13.035 23.305m8.288.111c5.948-3.93 8.33-9.48 9.497-16.066m-7.375 8.195l4.597 7.773"/>`),
		g.Group(children),
	)
}