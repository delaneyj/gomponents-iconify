package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ndax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.259 24L8.241 12.529v22.942L20.259 24z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.976 29.997v9.488L30.583 24L13.976 8.515v9.488"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.711 34.137V43.5L39.759 24L19.711 4.5v9.363"/>`),
		g.Group(children),
	)
}