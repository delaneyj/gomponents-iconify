package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 15.5h2m2 0h8m2 0h2M3.127 10.64c1.326.275 4.643.86 8.873.86c4.23 0 7.547-.585 8.873-.86m-17.746 0A16 16 0 0 0 4.5 4.155V3.5h15v.656a16 16 0 0 0 1.373 6.484m-17.746 0a16 16 0 0 1-1.314 2.39l-.313.47v7h2.25l.075-.12a4 4 0 0 1 3.392-1.88h9.566a4 4 0 0 1 3.392 1.88l.075.12h2.25v-7l-.313-.47a15.993 15.993 0 0 1-1.314-2.39"/>`),
		g.Group(children),
	)
}