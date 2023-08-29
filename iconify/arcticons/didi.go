package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Didi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.71 10.355H4.5v7.79a19.5 19.5 0 0 0 19.5 19.5h0a19.5 19.5 0 0 0 19.5-19.5"/>`),
		g.Group(children),
	)
}