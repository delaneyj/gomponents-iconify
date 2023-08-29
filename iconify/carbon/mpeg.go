package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mpeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 21h-4a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h4v2h-4v6h2v-2h-1v-2h3zm-8-8v-1.976h-6V21h6v-2h-4v-2h2v-2h-2v-2h4zm-10-2H9v10h2v-3h3a2.003 2.003 0 0 0 2-2v-3a2.002 2.002 0 0 0-2-2zm-3 5v-3h3l.001 3zm-4-5H5l-1.5 4L2 11H0v10h2v-7l1.5 4L5 14v7h2V11z"/>`),
		g.Group(children),
	)
}