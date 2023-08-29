package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.5 20a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Zm0-2.111a7.389 7.389 0 1 0 0-14.778a7.389 7.389 0 0 0 0 14.778Z" opacity=".2"/><path d="M10 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17ZM.5 10a9.5 9.5 0 1 1 19 0a9.5 9.5 0 0 1-19 0Z"/></g>`),
		g.Group(children),
	)
}