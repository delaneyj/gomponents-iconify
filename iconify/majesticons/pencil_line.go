package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.586 3a2 2 0 0 1 2.828 0L21 5.586a2 2 0 0 1 0 2.828l-12 12A2 2 0 0 1 7.586 21H5a2 2 0 0 1-2-2v-2.586A2 2 0 0 1 3.586 15l12-12zm-.172 3L18 8.586L19.586 7L17 4.414L15.414 6zm1.172 4L14 7.414l-9 9V19h2.586l9-9z"/></g>`),
		g.Group(children),
	)
}