package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlepencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 760L556 356L356 556l404 404q-116 64-248 64q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257q0 133-64 248zM256 256l34 112q57-21 78-78zm153 47q-12 38-40 66t-66 40l33 107l180-180z"/>`),
		g.Group(children),
	)
}