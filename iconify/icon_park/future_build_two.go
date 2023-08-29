package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44H44"/><ellipse cx="24.5" cy="7" rx="13.5" ry="3"/><path d="M16 9C16 9 20.1593 17.8828 21 24C22.0687 31.7763 20 44 20 44"/><path d="M32.2266 9C32.2266 9 28.0673 17.8828 27.2266 24C26.1578 31.7763 28.0006 44 28.0006 44"/></g>`),
		g.Group(children),
	)
}