package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulbOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h1m8-9v1m8 8h1M5.6 5.6l.7.7m12.1-.7l-.7.7m-6.611.783a5 5 0 0 1 5.826 5.84m-1.378 2.611A5.012 5.012 0 0 1 15 16a3.5 3.5 0 0 0-1 3a2 2 0 1 1-4 0a3.5 3.5 0 0 0-1-3a5 5 0 0 1-.528-7.544M9.7 17h4.6M3 3l18 18"/>`),
		g.Group(children),
	)
}