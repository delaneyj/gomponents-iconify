package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.437 20.938H5.562a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.873a2.5 2.5 0 0 1-2.5 2.5ZM5.562 4.065a1.5 1.5 0 0 0-1.5 1.5v12.873a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Z"/><circle cx="11.999" cy="12.002" r="1" fill="currentColor"/><circle cx="15.999" cy="12.002" r="1" fill="currentColor"/><circle cx="7.999" cy="12.002" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}