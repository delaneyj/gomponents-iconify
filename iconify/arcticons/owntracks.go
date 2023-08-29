package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Owntracks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.502a13.09 13.09 0 0 0-13.09 13.09c0 10.25 9.98 22.61 12.59 25.63a.59.59 0 0 0 1 0c2.55-3 12.59-15.38 12.59-25.63A13.09 13.09 0 0 0 24 4.502Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.42 21.74l5.62-6.693l5.54 6.65M24 15.047v28.404"/>`),
		g.Group(children),
	)
}