package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M4 0h.971v1.911H4zm7 0h1v1.906h-1z"/><path d="M15.976 4.959V2.456c0-.771-.606-1.398-1.354-1.398h-1.59v2.026H9.968V1.058H6.034v2.026H2.938V1.058H1.401c-.748 0-1.354.627-1.354 1.398v2.503h15.929zM.046 6.003v8.562c0 .772.606 1.399 1.354 1.399h13.221c.748 0 1.354-.627 1.354-1.399V6.003H.046z"/></g>`),
		g.Group(children),
	)
}