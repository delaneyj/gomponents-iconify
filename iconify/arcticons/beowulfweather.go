package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beowulfweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.149 23.716C10.71 12.371 25.12 11.311 26.724 21.688c3.489-2.662 7.874.279 7.05 4.917c8.227.074 7.03 9.678 2.226 9.925H10.026c-7.018.174-8.418-12.394 2.123-12.814l-.267-.173"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.1 16.008a10.668 10.668 0 1 1 15.636 14.26"/>`),
		g.Group(children),
	)
}