package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkRemoveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3V0h1v3H4ZM9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708L9.318.975ZM0 4h3v1H0V4Zm10.854.854l-6 6l-.708-.708l6-6l.708.708Zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172l.708.708ZM15 11h-3v-1h3v1Zm-5 4v-3h1v3h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}