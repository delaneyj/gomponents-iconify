package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.5H5a.5.5 0 0 0-.947-.224L4.5 3.5Zm0 8v.5a.5.5 0 0 0 .5-.5h-.5Zm-4 0l-.447-.224A.5.5 0 0 0 .5 12v-.5Zm10-8l.447-.224A.5.5 0 0 0 10 3.5h.5Zm0 8H10a.5.5 0 0 0 .5.5v-.5Zm4 0v.5a.5.5 0 0 0 .447-.724l-.447.224ZM4 3.5v8h1v-8H4Zm.5 7.5h-4v1h4v-1Zm-3.553.724l4-8l-.894-.448l-4 8l.894.448ZM10 3.5v8h1v-8h-1Zm.5 8.5h4v-1h-4v1Zm4.447-.724l-4-8l-.894.448l4 8l.894-.448ZM7 0v15h1V0H7Z"/>`),
		g.Group(children),
	)
}