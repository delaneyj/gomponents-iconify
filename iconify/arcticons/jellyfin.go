package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jellyfin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 20c-1.62 0-6.85 9.48-6.06 11.08s11.33 1.59 12.12 0S25.63 20 24 20Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5c-4.89 0-20.66 28.58-18.25 33.4s34.13 4.77 36.51 0S28.9 5.5 24 5.5Zm12 29.21c-1.56 3.13-22.35 3.17-23.93 0S20.8 12.83 24 12.83s13.52 18.76 12 21.88Z"/>`),
		g.Group(children),
	)
}