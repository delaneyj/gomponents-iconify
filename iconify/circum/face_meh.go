package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMeh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942Zm0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058Z"/><circle cx="9.001" cy="8.99" r="1.25" fill="currentColor"/><circle cx="15.001" cy="8.99" r="1.25" fill="currentColor"/><path fill="currentColor" d="M8.438 15.939h7.125a.5.5 0 0 0 0-1H8.438a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}