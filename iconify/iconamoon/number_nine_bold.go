package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.941 19.335a1.25 1.25 0 1 0 2.118 1.33l-2.118-1.33Zm7.345-6.998a1.25 1.25 0 1 0-2.118-1.33l2.118 1.33ZM8.25 9A3.75 3.75 0 0 1 12 5.25v-2.5A6.25 6.25 0 0 0 5.75 9h2.5ZM12 5.25A3.75 3.75 0 0 1 15.75 9h2.5A6.25 6.25 0 0 0 12 2.75v2.5ZM15.75 9A3.75 3.75 0 0 1 12 12.75v2.5A6.25 6.25 0 0 0 18.25 9h-2.5ZM12 12.75A3.75 3.75 0 0 1 8.25 9h-2.5A6.25 6.25 0 0 0 12 15.25v-2.5Zm.059 7.915l5.227-8.328l-2.118-1.33l-5.227 8.329l2.118 1.329Z"/>`),
		g.Group(children),
	)
}