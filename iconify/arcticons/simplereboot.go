package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplereboot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.909 42.699a16.486 16.486 0 0 1-7.92-25.796m6.689-5.109A16.483 16.483 0 1 1 24 43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.137 4.5l-4.458 7.298l7.299 4.458"/>`),
		g.Group(children),
	)
}