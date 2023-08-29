package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NtCalculadora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="37" height="37" x="5.5" y="5.5" rx="2" ry="2"/><path d="M28.727 15.426h9m-9 17.831h9m-26.471-3.935l8 7.937m-.039-8l-7.96 7.969m3.999-26.302v9m-4.5-4.5h9"/></g><circle cx="33.227" cy="30.265" r=".75" fill="currentColor"/><circle cx="33.227" cy="36.249" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}