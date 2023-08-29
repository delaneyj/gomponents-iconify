package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cineaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="26" x="4.5" y="13.014" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="25.527" height="17.701" x="8.286" y="17.164" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.258 13.014l3.05-4.028m-4.788 1.093l1.738 2.935m1.555 13l-9.527-5.5v11l9.527-5.5zm10.804-8.85h4.374m-4.374 3.35h4.374"/><circle cx="38.804" cy="26.248" r="2.187" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.804" cy="32.678" r="2.187" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}