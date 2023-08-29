package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.75 8.83a4.34 4.34 0 0 1 8.67 0m0 30.34a4.34 4.34 0 0 1-8.67 0m-9.21-25.05l10.84-8.67m5.41 6.77L17 20.88M29.42 8.83v30.34m-8.67 0V8.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.79 12.22a4.33 4.33 0 0 0-2.71-7.72H25a4.37 4.37 0 0 0-2.64 1m-10.82 8.62a4.33 4.33 0 0 0 2.71 7.71h.07a4.31 4.31 0 0 0 2.68-.95m3.75 9.27l8.67 6.93"/>`),
		g.Group(children),
	)
}