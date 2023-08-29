package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h6v-.45q0-.425-.238-.788T11.1 16.2q-.5-.225-1.012-.338T9 15.75q-.575 0-1.088.113T6.9 16.2q-.425.2-.663.563T6 17.55V18Zm8-1.5h4V15h-4v1.5ZM9 15q.625 0 1.063-.438T10.5 13.5q0-.625-.438-1.063T9 12q-.625 0-1.063.438T7.5 13.5q0 .625.438 1.063T9 15Zm5-1.5h4V12h-4v1.5ZM2 22V7h7V2h6v5h7v15H2Zm2-2h16V9h-5v2H9V9H4v11Zm7-11h2V4h-2v5Zm1 5.5Z"/>`),
		g.Group(children),
	)
}