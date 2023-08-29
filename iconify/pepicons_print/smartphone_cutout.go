package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneCutout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 2.5A1.5 1.5 0 0 1 7.5 1h8A1.5 1.5 0 0 1 17 2.5v16a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 6 18.5v-16Z" opacity=".2"/><path d="M4 2A1.5 1.5 0 0 1 5.5.5h9A1.5 1.5 0 0 1 16 2v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 4 18V2Zm1.5-.5A.5.5 0 0 0 5 2v16a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V2a.5.5 0 0 0-.5-.5h-9Z"/><path d="M7.5 3a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}