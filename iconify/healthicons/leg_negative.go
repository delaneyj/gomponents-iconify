package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsLegNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM21.33 4.79l-3.196 14.882a5.996 5.996 0 0 0-.134 1.26v2.606a6 6 0 0 0 .301 1.876L22.5 38.167L16 40.795h-3v1.154a2 2 0 0 0 2 2h3.31c.127 0 .252-.012.376-.036l3.982-.761a3.878 3.878 0 0 1 2.532.376a2.908 2.908 0 0 0 4.032-1.441l.055-.132a3.04 3.04 0 0 0-.444-3.104L27 36.59l.397-9.181a8.001 8.001 0 0 0-1.116-4.434L25 20.82s7.266-8.115 8.814-15.702c.12-.59-.347-1.118-.949-1.118H22.308a1 1 0 0 0-.978.79Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsLegNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}