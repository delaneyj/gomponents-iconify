package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.175 2.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H6.675a.5.5 0 0 1-.5-.5V2.5Zm1 .5v2.875h5.643V3H7.175Z" clip-rule="evenodd"/><path d="M4.5 5v12h11V5h-2V4h2a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h2v1h-2Z"/><path fill-rule="evenodd" d="M12.284 9.088a.5.5 0 0 1 .128.696l-1.767 2.564a1 1 0 0 1-1.437.222l-1.515-1.175a.5.5 0 1 1 .614-.79L9.82 11.78l1.767-2.564a.5.5 0 0 1 .696-.128Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}