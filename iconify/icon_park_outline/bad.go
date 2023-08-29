package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9.948 28.807C13.869 37.681 21.933 41.19 28 44c2.632 1.22 3.328-3.717 2.277-6.69c-.884-2.502-2.651-5.305-2.651-5.305h7.106a3.5 3.5 0 1 0 0-7.001h1.767a3.5 3.5 0 1 0 0-7.002h-3.535a3.5 3.5 0 0 0 0-7h-2.652A3.5 3.5 0 0 0 30.315 4H19.67c-3.094 0-7.071 1.803-9.723 6.804c-2.542 4.794-2.652 12.002 0 18.003Z"/>`),
		g.Group(children),
	)
}