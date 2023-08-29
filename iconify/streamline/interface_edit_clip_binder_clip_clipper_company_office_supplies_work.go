package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditClipBinderClipClipperCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 6h1a1 1 0 0 1 1 1.13l-.73 5.5a1 1 0 0 1-1 .87H2.23a1 1 0 0 1-1-.87l-.72-5.5A1 1 0 0 1 1.5 6h1"/><path d="M10.46 13.5L8.33 4h.42a1.75 1.75 0 0 0 0-3.5h-3.5a1.75 1.75 0 0 0 0 3.5h.41l-2.12 9.5"/></g>`),
		g.Group(children),
	)
}