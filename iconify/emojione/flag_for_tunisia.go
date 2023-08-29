package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTunisia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ed4c5c"/><circle cx="32" cy="32" r="17.3" fill="#fff"/><circle cx="32" cy="32" r="13" fill="#ed4c5c"/><circle cx="35.5" cy="32" r="10.4" fill="#fff"/><path fill="#ed4c5c" d="m38.4 32l3.4-4.6l-5.4 1.8l-3.3-4.6v5.7L27.7 32l5.4 1.7v5.7l3.3-4.6l5.4 1.8z"/>`),
		g.Group(children),
	)
}