package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.021 4H19a1 1 0 1 1 0 2h-4.246l-3.428 12H15a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2h4.246l3.428-12H9a1 1 0 0 1 0-2h5.021z"/>`),
		g.Group(children),
	)
}