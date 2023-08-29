package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.874 26.237L35 24l3.874-2.237a5.5 5.5 0 0 0-5.5-9.526L29.5 14.474V10a5.5 5.5 0 0 0-11 0v4.474l-3.874-2.237a5.5 5.5 0 0 0-5.5 9.526L13 24l-3.874 2.237a5.5 5.5 0 0 0 5.5 9.526l3.874-2.237V38a5.5 5.5 0 0 0 11 0v-4.474l3.874 2.237a5.5 5.5 0 0 0 5.5-9.526ZM18.5 14.474L35 24m-16.5 9.526L35 24"/>`),
		g.Group(children),
	)
}