package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneCutoutCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSmartphoneCutoutCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 5a1.5 1.5 0 0 1 1.5-1.5h9A1.5 1.5 0 0 1 19 5v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 7 21V5Zm1.5-.5A.5.5 0 0 0 8 5v16a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5h-9Z"/><path d="M10.5 6a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSmartphoneCutoutCircleFilled0)"/></g>`),
		g.Group(children),
	)
}