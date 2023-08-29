package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1 7.01C1 5.898 1.898 5 3.01 5H29c1.109 0 2 .904 2 2.01V25c0 1.109-.904 2-2.01 2H3.01A2.007 2.007 0 0 1 1 24.99V7.01Zm2-.005v.497l12.326 7.854c.413.262.935.262 1.348 0L29 7.51v-.508L28.996 7H3.005l-.003.002l-.001.002v.001Zm26 2.877l-10.047 6.394L29 22.671V9.88Zm-12.01 7.516a3.253 3.253 0 0 1-1.98 0L3.08 25h25.855l-11.944-7.602Zm-3.941-1.122L3 9.873V22.68l10.049-6.404Z"/>`),
		g.Group(children),
	)
}