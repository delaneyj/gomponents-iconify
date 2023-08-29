package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M24 100h28v72H32a4 4 0 0 0 0 8h192a4 4 0 0 0 0-8h-20v-72h28a4 4 0 0 0 2.1-7.41l-104-64a4 4 0 0 0-4.2 0l-104 64A4 4 0 0 0 24 100Zm36 0h40v72H60Zm88 0v72h-40v-72Zm48 72h-40v-72h40ZM128 36.7L217.87 92H38.13ZM244 208a4 4 0 0 1-4 4H16a4 4 0 0 1 0-8h224a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}