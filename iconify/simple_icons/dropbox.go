package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 1.807L0 5.629l6 3.822l6.001-3.822L6 1.807zm12 0l-6 3.822l6 3.822l6-3.822l-6-3.822zM0 13.274l6 3.822l6.001-3.822L6 9.452l-6 3.822zm18-3.822l-6 3.822l6 3.822l6-3.822l-6-3.822zM6 18.371l6.001 3.822l6-3.822l-6-3.822L6 18.371z"/>`),
		g.Group(children),
	)
}