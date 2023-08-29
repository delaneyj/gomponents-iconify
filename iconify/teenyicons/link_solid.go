package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708L9.318.975Zm1.536 3.879l-6 6l-.708-.708l6-6l.708.708Zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172l.708.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}