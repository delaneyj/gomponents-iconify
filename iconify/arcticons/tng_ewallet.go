package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TngEwallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.792 35.074c-8.995 0-16.287-7.292-16.287-16.287c0-7.893 5.616-14.472 13.069-15.966A21.54 21.54 0 0 0 24 2.501C12.126 2.5 2.5 12.125 2.5 24S12.126 45.5 24 45.5c11.215 0 20.414-8.59 21.401-19.548c-2.652 5.398-8.189 9.122-14.61 9.122Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.288 26.95c-2.006 8.902-9.96 15.55-19.467 15.55c-11.021 0-19.955-8.934-19.955-19.955S14.8 2.59 25.82 2.59"/>`),
		g.Group(children),
	)
}