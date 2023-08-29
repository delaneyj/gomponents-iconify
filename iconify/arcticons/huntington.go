package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Huntington(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.482 4.5v11.592m-5.936-7.847v31.51M8.61 11.989v23.943m5.936-9.948l11.872-6.854V4.5m1.1 39V31.908m5.936 7.847V8.245m5.936 27.766V12.068m-5.936 9.948L21.582 28.87V43.5"/>`),
		g.Group(children),
	)
}