package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayStoreAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.3 5.52a7 7 0 0 1 2 .76L32.52 18l-5.87 6L7.68 5.9a2.46 2.46 0 0 1 1.62-.38Zm-1.62.38l19 18.1L7.7 42.07c-.7-.56-1.07-1.69-1.07-3.36V9.29c0-1.67.36-2.82 1-3.38ZM32.52 18l7 3.87c2.49 1.38 2.49 2.84 0 4.22l-7 3.87l-5.87-6l5.87-6Zm0 12L11.34 41.72c-1.16.64-2.72 1.19-3.64.35L26.65 24Z"/>`),
		g.Group(children),
	)
}