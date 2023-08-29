package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#1F4C9F"/><path fill="#FFF" d="M6 19.25L16 6.5l10 12.75v4.25L16 10.75L6 23.5v-4.25zm10.5 1.25a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/></g>`),
		g.Group(children),
	)
}