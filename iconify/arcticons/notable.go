package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.65 9.75c-4.08 3.68-11.36 9.77-12.78 11.54a7.49 7.49 0 0 0-1.3 1.89c3.5 1.4 7.49-.11 8.88 3.89a4 4 0 0 1-.57 3A5.25 5.25 0 0 1 28.66 32A4.7 4.7 0 0 1 23 31c-1.34-1.55-1.47-4.23-.47-7.86A4.47 4.47 0 0 1 19.48 19c-5.14 4.41-11.79 4.5-14 4.4v17.2a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v15.9m24.56-14h5.85a1 1 0 0 1 1 1v5.84m-25.32-1v-4.87a1 1 0 0 1 1-1h4.87m0 29.21h-4.92a1 1 0 0 1-1-1v-4.82m25.32 0v4.87a1 1 0 0 1-1 1H31"/>`),
		g.Group(children),
	)
}