package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9 7c0-.507.16-1.289.612-1.916C10.026 4.508 10.726 4 12 4c1.274 0 1.974.508 2.389 1.084c.45.627.611 1.41.611 1.916a1 1 0 1 0 2 0c0-.827-.24-2.044-.988-3.084C15.226 2.825 13.926 2 12 2c-1.926 0-3.226.825-4.012 1.916C7.24 4.956 7 6.173 7 7v3H6a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3H9V7zm4 8a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}