package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20 12.527v8.946a4.5 4.5 0 1 0 0-8.945Zm-13.243-.464a.5.5 0 0 1 .68.194L9.5 15.97l2.063-3.713a.5.5 0 0 1 .874.486L10.072 17H12a.5.5 0 0 1 0 1h-2v1h2a.5.5 0 0 1 0 1h-2v1.5a.5.5 0 0 1-1 0V20H7a.5.5 0 0 1 0-1h2v-1H7a.5.5 0 0 1 0-1h1.928l-2.365-4.257a.5.5 0 0 1 .194-.68Z"/><path d="M4.5 7a3 3 0 0 0-3 3v18A2.5 2.5 0 0 0 4 30.5h10V25H4.5a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1H15v1H5.5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1H15v6.5h4V24h7.5a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1H19V9h8.5a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H20v5.5h8a2.5 2.5 0 0 0 2.5-2.5V10a3 3 0 0 0-3-3h-23ZM19 11h7.5v12H19V11ZM5.5 23V11H15v12H5.5Z"/></g>`),
		g.Group(children),
	)
}