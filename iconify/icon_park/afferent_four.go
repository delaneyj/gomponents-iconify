package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfferentFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21 5L10 5C8.89543 5 8 5.89543 8 7L8 41C8 42.1046 8.89543 43 10 43L38 43C39.1046 43 40 42.1046 40 41L40 24.75"/><path d="M33 24H21V12"/><path d="M21.0001 23.9998L39 6"/></g>`),
		g.Group(children),
	)
}