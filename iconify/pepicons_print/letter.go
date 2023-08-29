package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Letter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><rect width="15" height="12" x="4.5" y="5.5" opacity=".2" rx="1.5"/><path fill-rule="evenodd" d="M17 4H3a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5v-11A.5.5 0 0 0 17 4ZM3.5 15V5h13v10h-13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m17.324 4.88l-7.045 6a.5.5 0 0 1-.65-.001l-6.956-6A.5.5 0 0 1 3 4h14a.5.5 0 0 1 .324.88ZM15.642 5H4.345l5.612 4.841L15.642 5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}