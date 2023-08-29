package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 30l-2.139-1.013A5.022 5.022 0 0 1 18 24.467V18h10v6.468a5.022 5.022 0 0 1-2.861 4.52Zm-3-10v4.468a3.012 3.012 0 0 0 1.717 2.71l1.283.608l1.283-.607A3.012 3.012 0 0 0 26 24.468V20Z"/><path fill="currentColor" d="M16 28H6V4h8v6a2.006 2.006 0 0 0 2 2h6v3h2v-5a.91.91 0 0 0-.3-.7l-7-7A.909.909 0 0 0 16 2H6a2.006 2.006 0 0 0-2 2v24a2.006 2.006 0 0 0 2 2h10Zm0-23.6l5.6 5.6H16Z"/>`),
		g.Group(children),
	)
}