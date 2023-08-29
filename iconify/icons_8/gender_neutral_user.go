package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderNeutralUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5c-3.854 0-7 3.146-7 7c0 2.41 1.23 4.552 3.094 5.813C8.527 19.343 6 22.88 6 27h2c0-4.43 3.57-8 8-8s8 3.57 8 8h2c0-4.12-2.527-7.658-6.094-9.188A7.02 7.02 0 0 0 23 12c0-3.854-3.146-7-7-7zm0 2c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5z"/>`),
		g.Group(children),
	)
}