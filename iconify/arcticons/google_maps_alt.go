package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMapsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v6.7l28.3 28.3h6.75a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95Zm22.71 3.89a8.47 8.47 0 0 1 8.45 8.45h0c0 4.92-4 8.75-8.45 13.78c-4.26-4.82-8.1-8.54-8.42-13.16v-.62a8.45 8.45 0 0 1 8.45-8.45Zm0 4.39a4.06 4.06 0 0 0-4.06 4.06h0a4.06 4.06 0 0 0 4.06 4.06h0a4.06 4.06 0 0 0 0-8.12ZM5.5 27.18v13.37a2 2 0 0 0 2 2h13.32Z"/>`),
		g.Group(children),
	)
}