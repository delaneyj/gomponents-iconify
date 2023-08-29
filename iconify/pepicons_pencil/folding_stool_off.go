package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M4 11a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1Zm2-9.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V2a.5.5 0 0 1 .5-.5Z"/><path d="M6 6.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V7a.5.5 0 0 1 .5-.5Zm8 0a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V7a.5.5 0 0 1 .5-.5Z"/><path d="M14 1.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V2a.5.5 0 0 1 .5-.5Zm1.413 16.782a.5.5 0 0 1-.695.13l-9.5-6.5a.5.5 0 0 1 .564-.825l9.5 6.5a.5.5 0 0 1 .13.695Z"/><path d="M4.587 18.282a.5.5 0 0 0 .695.13l9.5-6.5a.5.5 0 0 0-.564-.825l-9.5 6.5a.5.5 0 0 0-.13.695ZM14 4H6V3h8v1Zm0 3H6V6h8v1ZM2.5 8a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5Zm11 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}