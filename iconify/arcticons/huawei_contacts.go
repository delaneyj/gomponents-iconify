package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiContacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="33.712" height="30.958" x="9.788" y="8.521" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.051"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16.288h9.695M4.5 31.105h9.695"/><circle cx="27.553" cy="20.061" r="6.481" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.605 36.173l-9.912-9.633l-10.028 9.633"/>`),
		g.Group(children),
	)
}