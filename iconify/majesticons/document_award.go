package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentAward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 21h14c1.126 0 1.926-.491 2.412-1.166c.448-.623.588-1.34.588-1.834v-6a1 1 0 0 0-1-1h-3V6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v12c0 .493.14 1.211.588 1.834C3.074 20.51 3.874 21 5 21zm14-8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1zM9 10a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-3 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}