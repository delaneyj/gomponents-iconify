package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.32 29.35A12.38 12.38 0 0 0 36 27.7a14.21 14.21 0 0 0 3.55-9.27A13.93 13.93 0 0 0 25.62 4.5H14m-5.55 5.57v27.86a5.57 5.57 0 1 0 11.14 0V26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.2 20.47a5.57 5.57 0 0 0-4.56 8.77L29 41.13a5.57 5.57 0 0 0 9.12-6.38l-8.36-11.91a5.55 5.55 0 0 0-4.56-2.38Z"/><circle cx="14.02" cy="10.07" r="5.57" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}