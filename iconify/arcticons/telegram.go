package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telegram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.83 8.48c1.14 0 2 1 1.54 2.86l-5.58 26.3c-.39 1.87-1.52 2.32-3.08 1.45L20.4 29.26a.4.4 0 0 1 0-.65l15.37-13.88c.7-.62-.15-.92-1.07-.36L15.41 26.54a.46.46 0 0 1-.4.05L6.82 24C5 23.47 5 22.22 7.23 21.33L40 8.69a2.16 2.16 0 0 1 .83-.21Z"/>`),
		g.Group(children),
	)
}