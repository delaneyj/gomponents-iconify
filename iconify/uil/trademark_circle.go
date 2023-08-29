package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrademarkCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9H7a1 1 0 0 0 0 2h.5v3a1 1 0 0 0 2 0v-3h.5a1 1 0 0 0 0-2Zm7.38.08a1 1 0 0 0-1.09.21L15 10.59l-1.29-1.3a1 1 0 0 0-1.09-.21A1 1 0 0 0 12 10v4a1 1 0 0 0 2 0v-1.59l.29.3a1 1 0 0 0 1.42 0l.29-.3V14a1 1 0 0 0 2 0v-4a1 1 0 0 0-.62-.92ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}