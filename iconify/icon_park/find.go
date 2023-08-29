package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Find(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M4 7H44"/><path stroke-linecap="round" d="M4 23H15"/><path stroke-linecap="round" d="M4 39H15"/><path fill="#2F88FF" d="M31.5 34C36.1944 34 40 30.1944 40 25.5C40 20.8056 36.1944 17 31.5 17C26.8056 17 23 20.8056 23 25.5C23 30.1944 26.8056 34 31.5 34Z"/><path stroke-linecap="round" d="M37 32L44 39.0505"/></g>`),
		g.Group(children),
	)
}