package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11 11h-1v-1h1v1Zm-3 0h1v-1H8v1Zm5 0h-1v-1h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M3 6V3.5a3.5 3.5 0 1 1 7 0V6h1.5A1.5 1.5 0 0 1 13 7.5v.55a2.5 2.5 0 0 1 0 4.9v.55a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 0 13.5v-6A1.5 1.5 0 0 1 1.5 6H3Zm1-2.5a2.5 2.5 0 0 1 5 0V6H4V3.5ZM8.5 9a1.5 1.5 0 1 0 0 3h4a1.5 1.5 0 0 0 0-3h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}