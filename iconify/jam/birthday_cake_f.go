package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirthdayCakeF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 15.693l-5.141 1.282l-4.82-.996l-4.935.996L0 15.687V15a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v.693zm0 2.063V22H0v-4.25l5.057 1.275l4.978-1.004l4.859 1.004L20 17.756zM7 9h6v2H7V9zm3-.6a3 3 0 0 1-3-3C7 4.295 8 2.495 10 0c2 2.495 3 4.295 3 5.4a3 3 0 0 1-3 3z"/>`),
		g.Group(children),
	)
}