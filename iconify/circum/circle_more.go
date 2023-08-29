package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12.001" cy="12" r="1" fill="currentColor"/><circle cx="16.001" cy="12" r="1" fill="currentColor"/><circle cx="8.001" cy="12" r="1" fill="currentColor"/><path fill="currentColor" d="M12 21.932A9.934 9.934 0 1 1 21.934 12A9.944 9.944 0 0 1 12 21.932Zm0-18.867A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.065Z"/>`),
		g.Group(children),
	)
}