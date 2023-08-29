package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopperIudNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCopperIudNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm18.868 6H23v9.257l2 1V6h4.132l-1-2H25a1 1 0 0 0-1 1a1 1 0 0 0-1-1h-3.132l-1 2Zm-1.236-2h-1.764l-1 2h1.764l1-2ZM11 4h2.632l-1 2H11a1 1 0 1 1 0-2Zm14 14.493v1.764l-2-1v-1.764l2 1Zm-2 3v1.764l2 1v-1.764l-2-1Zm0 14.633V25.493l2 1v9.633A4.002 4.002 0 0 1 24 44a4 4 0 0 1-1-7.874ZM24 38a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm6.368-34h1.764l1 2h-1.764l-1-2ZM37 4h-2.632l1 2H37a1 1 0 1 0 0-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCopperIudNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}