package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="27.121" cy="17.107" r="12.607" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.975 29.662L14.73 42.234a3.68 3.68 0 0 1-5.965-4.255l7.78-14.007m15.777-12.13a7.204 7.204 0 0 1-.026 9.982m-14.015 9.92l-2.009 3.042"/>`),
		g.Group(children),
	)
}