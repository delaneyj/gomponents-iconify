package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zM28.139 45l-3.084-9.1h-8.111L13.861 45H10l8.805-26h4.389L32 45h-3.861zM52 37.467c0 4.143-3.064 7.516-6.832 7.516L36 45V19h10c3.768 0 6 3.391 6 7.535c0 2.152-.828 4.094-2.15 5.465c1.322 1.371 2.15 3.313 2.15 5.467z"/><path fill="currentColor" d="M45.168 34.051h-5.623v6.832h5.623c1.713 0 3.105-1.533 3.105-3.416c0-1.885-1.392-3.416-3.105-3.416M18.266 32h5.47L21 23.926zm26.902-8.883h-5.623v6.834h5.623c1.713 0 3.105-1.533 3.105-3.416s-1.392-3.418-3.105-3.418"/>`),
		g.Group(children),
	)
}