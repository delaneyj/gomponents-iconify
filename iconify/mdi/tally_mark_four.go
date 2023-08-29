package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TallyMarkFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19H5V5h2v14m4-14H9v14h2V5m4 0h-2v14h2V5m4 0h-2v14h2V5Z"/>`),
		g.Group(children),
	)
}