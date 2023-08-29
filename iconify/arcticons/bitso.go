package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.8c5.1 0 9.2 4.1 9.2 9.2v19.4c7.3-3.5 12.3-10.8 12.3-19.4c0-11.9-9.6-21.5-21.5-21.5c-1.3 0-2.6.1-3.9.4v11.9H24Zm0 18.4c-5.1 0-9.2-4.1-9.2-9.2V2.5H2.5V24c0 11.9 9.6 21.5 21.5 21.5c1.3 0 2.6-.1 3.9-.4V33.2H24Z"/>`),
		g.Group(children),
	)
}