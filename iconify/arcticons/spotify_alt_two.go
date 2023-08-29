package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpotifyAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm-4.89 7.76a4.37 4.37 0 0 1 1.75.54a37.37 37.37 0 0 1 16.33 16.34c1.21 2.25.25 3.89-.89 4.39s-3.31.44-4-.93a35.84 35.84 0 0 0-14.9-14.86c-1.37-.73-1.44-2.9-.93-4a2.52 2.52 0 0 1 2.14-1.44a2.78 2.78 0 0 1 .5-.04ZM15.18 19a1.83 1.83 0 0 1 1 .22a32.06 32.06 0 0 1 12.57 12.6a1.85 1.85 0 0 1-.84 2.58a2.07 2.07 0 0 1-2.69-.63a28.19 28.19 0 0 0-11-11a2.07 2.07 0 0 1-.63-2.69A1.77 1.77 0 0 1 15.18 19Zm-2.93 7.48a1.46 1.46 0 0 1 .76.22a24 24 0 0 1 8.26 8.3a1.32 1.32 0 0 1-.45 1.91a1.66 1.66 0 0 1-2.06-.39a21.25 21.25 0 0 0-7.27-7.27a1.67 1.67 0 0 1-.39-2.06a1.22 1.22 0 0 1 1.15-.68Z"/>`),
		g.Group(children),
	)
}