package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneHomeButtonCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSmartphoneHomeButtonCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M17.5 3.5h-9A1.5 1.5 0 0 0 7 5v16a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 19 21V5a1.5 1.5 0 0 0-1.5-1.5ZM8 5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v16a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 8 21V5Z" clip-rule="evenodd"/><path d="M13 21a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSmartphoneHomeButtonCircleFilled0)"/></g>`),
		g.Group(children),
	)
}