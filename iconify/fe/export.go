package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feExport0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feExport1" fill="currentColor"><path id="feExport2" d="M13 5.828V17h-2V5.828L7.757 9.071L6.343 7.657L12 2l5.657 5.657l-1.414 1.414L13 5.828ZM4 16h2v4h12v-4h2v4c0 1.1-.9 2-2 2H6c-1.1 0-2-.963-2-2v-4Z"/></g></g>`),
		g.Group(children),
	)
}