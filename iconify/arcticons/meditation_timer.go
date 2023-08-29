package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeditationTimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.8 39.232a21.498 21.498 0 0 1 11.943-36.45m.23-.034a21.498 21.498 0 0 1 9.111 41.902m-17.508-9.195a16.156 16.156 0 0 1-1.599-20.985m.105-.14a16.156 16.156 0 0 1 29.063 9.073a18.587 18.587 0 0 1-.704 5.42a19.13 19.13 0 0 1-2.012 4.92a24.552 24.552 0 0 1-4.223 5.589a14.367 14.367 0 0 1-3.827 2.485m-16.891-6.451a2.887 2.887 0 0 1 .808 1.557a2.584 2.584 0 0 1-2.229 3.038a2.843 2.843 0 0 1-2.349-.827m16.587 4.864c0 1.388-2.652 1.55-3.695 1.35c-1.418-.272-3.313-.666-3.14-1.872c.182-1.269 2.72-1.227 3.698-1.13c.907.09 3.137.406 3.137 1.652Zm4.074-2.18c-1.13.565-2.01.697-1.811 1.832c.119.68.46 1.598 2.516 1"/>`),
		g.Group(children),
	)
}