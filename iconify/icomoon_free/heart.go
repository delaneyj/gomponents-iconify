package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.8 1c-1.682 0-3.129 1.368-3.799 2.797C7.33 2.368 5.883 1 4.201 1a4.202 4.202 0 0 0-4.2 4.2c0 4.716 4.758 5.953 8 10.616c3.065-4.634 8-6.05 8-10.616c0-2.319-1.882-4.2-4.2-4.2z"/>`),
		g.Group(children),
	)
}