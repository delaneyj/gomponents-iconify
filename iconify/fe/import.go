package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feImport0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feImport1" fill="currentColor"><path id="feImport2" d="m13 13.175l3.243-3.242l1.414 1.414L12 17.004l-5.657-5.657l1.414-1.414L11 13.175V2h2v11.175ZM4 16h2v4h12v-4h2v4c0 1.1-.9 2-2 2H6c-1.1 0-2-.963-2-2v-4Z"/></g></g>`),
		g.Group(children),
	)
}