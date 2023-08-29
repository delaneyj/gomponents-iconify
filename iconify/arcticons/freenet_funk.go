package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreenetFunk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.26 38.5V25.287L23.014 38.5V25.287M25.152 9.5v8.754c0 2.477 1.982 4.459 4.294 4.459s4.294-1.982 4.294-4.46V9.5m-19.476 6.607h4.295m-4.295 6.606V9.5h6.607m5.932 15.787V38.5m0-4.625l6.937-8.588m0 13.213l-5.285-6.607"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}