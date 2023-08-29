package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garminconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45 15.38a22.15 22.15 0 0 0-8.4-10a22.29 22.29 0 0 0-26.8 2.4A21.74 21.74 0 0 0 3 20.88h7.4a14.72 14.72 0 0 1 10.3-11.1c7.1-1.9 12.6.7 16.1 5.6Zm0 17.2a22 22 0 0 1-39.4 2.2A19.48 19.48 0 0 1 3 27h7.4a14.66 14.66 0 0 0 26.3 5.5Z"/>`),
		g.Group(children),
	)
}