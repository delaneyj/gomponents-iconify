package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneNotchCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSmartphoneNotchCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M8.5 3.5A1.5 1.5 0 0 0 7 5v16a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 19 21V5a1.5 1.5 0 0 0-1.5-1.5h-9ZM8 5a.5.5 0 0 1 .5-.5H10v1a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-1h1.5a.5.5 0 0 1 .5.5v16a.5.5 0 0 1-.5.5h-9A.5.5 0 0 1 8 21V5Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSmartphoneNotchCircleFilled0)"/></g>`),
		g.Group(children),
	)
}