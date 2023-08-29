package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindPowerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23q0-.825.588-1.413T12 21v-5.725q.225.1.475.15t.525.05q1.05 0 1.763-.712T15.474 13v-.225l4.025.975l3.55 6.35l-2.825 2.825l-6.225-6.2V21q.825 0 1.413.588T16 23h-6Zm-7-2v-2h5v2H3Zm8.375-6.125L8.05 17L1 15v-4h10.525q-.475.35-.738.863T10.525 13q0 .575.225 1.05t.625.825ZM13 14.5q-.625 0-1.063-.438T11.5 13q0-.625.438-1.063T13 11.5q.625 0 1.063.438T14.5 13q0 .625-.438 1.063T13 14.5Zm2.3-2.45q-.275-.675-.888-1.1T13 10.525q-.275 0-.525.05t-.475.15V5.9L17.35.925l3.375 2.1L15.3 12.05ZM1 9V7h5v2H1Zm3-4V3h6v2H4Z"/>`),
		g.Group(children),
	)
}