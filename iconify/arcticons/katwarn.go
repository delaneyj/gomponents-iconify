package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Katwarn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.876 39.344a5.434 5.434 0 0 0-10.867 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.82 31.222a13.5 13.5 0 0 0-24.281 8.122"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.864A21.952 21.952 0 0 0 9.82 39.417"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.45 10.606A30.872 30.872 0 0 0 4.5 26.363"/>`),
		g.Group(children),
	)
}