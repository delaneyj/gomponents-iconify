package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.39 16.08l-6.69-2.87l-2.54-4.34l-.08-.12A2 2 0 0 0 18.52 8h-8a2 2 0 0 0-1.67.89L5.46 14H3a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h2.14a4 4 0 0 0 7.72 0h6.28a4 4 0 0 0 7.72 0H29a1 1 0 0 0 1-1v-7a1 1 0 0 0-.61-.92ZM9 26a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm14 0a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm5-3h-1.14a4 4 0 0 0-7.72 0h-6.28a4 4 0 0 0-7.72 0H4v-7h2a1 1 0 0 0 .83-.45L10.54 10h8l2.63 4.5a1 1 0 0 0 .47.42L28 17.66Z"/>`),
		g.Group(children),
	)
}