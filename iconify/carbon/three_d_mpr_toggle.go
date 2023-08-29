package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDMprToggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M11 2H2v9h2V4h7V2z" fill="currentColor"/><path d="M2 21v9h9v-2H4v-7H2z" fill="currentColor"/><path d="M30 11V2h-9v2h7v7h2z" fill="currentColor"/><path d="M21 30h9v-9h-2v7h-7v2z" fill="currentColor"/><path d="M25.49 10.13l-9-5a1 1 0 0 0-1 0l-9 5A1 1 0 0 0 6 11v10a1 1 0 0 0 .51.87l9 5a1 1 0 0 0 1 0l9-5A1 1 0 0 0 26 21V11a1 1 0 0 0-.51-.87zM16 7.14L22.94 11L16 14.86L9.06 11zM8 12.7l7 3.89v7.71l-7-3.89zm9 11.6v-7.71l7-3.89v7.71z" fill="currentColor"/>`),
		g.Group(children),
	)
}