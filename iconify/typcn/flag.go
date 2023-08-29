package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.383 4.318a1 1 0 0 0-1.09.217a3.248 3.248 0 0 1-4.586 0a5.25 5.25 0 0 0-7.414 0A.997.997 0 0 0 5 5.242v13a1 1 0 1 0 2 0v-4.553a3.248 3.248 0 0 1 4.293.26a5.25 5.25 0 0 0 7.414 0a1 1 0 0 0 .293-.707v-8a1 1 0 0 0-.617-.924z"/>`),
		g.Group(children),
	)
}