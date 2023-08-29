package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.059 4.665a1.25 1.25 0 0 0-2.118-1.33l2.118 1.33Zm-7.345 6.998a1.25 1.25 0 0 0 2.118 1.33l-2.118-1.33ZM15.75 15A3.75 3.75 0 0 1 12 18.75v2.5A6.25 6.25 0 0 0 18.25 15h-2.5ZM12 18.75A3.75 3.75 0 0 1 8.25 15h-2.5A6.25 6.25 0 0 0 12 21.25v-2.5ZM8.25 15A3.75 3.75 0 0 1 12 11.25v-2.5A6.25 6.25 0 0 0 5.75 15h2.5ZM12 11.25A3.75 3.75 0 0 1 15.75 15h2.5A6.25 6.25 0 0 0 12 8.75v2.5Zm-.059-7.915l-5.227 8.328l2.118 1.33l5.227-8.329l-2.118-1.328Z"/>`),
		g.Group(children),
	)
}