package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataEnrichment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 13h2v2h-2zm-4.222-8.203l1.414-1.415l1.414 1.415l-1.414 1.414zM15 0h2v2h-2zM6.808 6.234L5.394 4.819l1.414-1.414L8.222 4.82zM2 13h2v2H2zm11 17h6v2h-6zm-2-4h10v2H11zm5-22C10.5 4 6 8.5 6 14c0 4.4 2 6.3 3.5 7.6c1 .9 1.5 1.5 1.5 2.4h2c0-1.8-1.1-2.9-2.2-3.9C9.4 18.9 8 17.5 8 14c0-4.4 3.6-8 8-8s8 3.6 8 8c0 3.5-1.4 4.9-2.8 6.1c-1.1 1-2.2 2-2.2 3.9h2c0-.9.5-1.5 1.5-2.4C24 20.3 26 18.4 26 14c0-5.5-4.5-10-10-10z"/>`),
		g.Group(children),
	)
}