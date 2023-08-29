package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M3.25 3a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1V3Z"/><path fill-rule="evenodd" d="M5.25 4v14h11V4h-11Zm-1-2a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1h-13Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M1.75 2A1.5 1.5 0 0 1 3.25.5h13a1.5 1.5 0 0 1 1.5 1.5v16a1.5 1.5 0 0 1-1.5 1.5h-13a1.5 1.5 0 0 1-1.5-1.5V2Zm1.5-.5a.5.5 0 0 0-.5.5v16a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V2a.5.5 0 0 0-.5-.5h-13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.75 17a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}