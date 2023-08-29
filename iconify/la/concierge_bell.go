package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConciergeBell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 6v2h6V6zm3 3C9.703 9 4.574 13.84 4.062 20H2v5h28v-5h-2.063C27.427 13.84 22.297 9 16 9zm0 2a9.927 9.927 0 0 1 9.938 9H6.063c.492-5.086 4.71-9 9.937-9zM4 22h24v1H4z"/>`),
		g.Group(children),
	)
}