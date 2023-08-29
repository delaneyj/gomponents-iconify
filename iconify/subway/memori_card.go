package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoriCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416.7 0H118L32.7 85.3v384c0 23.5 19.1 42.7 42.7 42.7h341.3c23.6 0 42.7-19.1 42.7-42.7V42.7c0-23.6-19.2-42.7-42.7-42.7zm-256 128H118V42.7h42.7V128zm64 0H182V42.7h42.7V128zm64 0H246V42.7h42.7V128zm64 0H310V42.7h42.7V128zm64 0H374V42.7h42.7V128z"/>`),
		g.Group(children),
	)
}