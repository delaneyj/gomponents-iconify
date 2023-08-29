package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastlyrics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.895 14.269h13.069m-14.262 6.84h13.069M4.5 33.731h13.069M43.5 14.269l-3.537 19.462h-8.57m-8.254 0l3.537-19.462h11.55m-12.793 6.84h11.56"/>`),
		g.Group(children),
	)
}