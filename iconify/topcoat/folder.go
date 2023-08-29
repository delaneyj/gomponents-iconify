package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M.5 7.5v26c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-21c0-2.46-.471-3-3-3h-17v-2c0-2.5-.52-3-3-3h-14c-2.48 0-3 .43-3 3z"/>`),
		g.Group(children),
	)
}