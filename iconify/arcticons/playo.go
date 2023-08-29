package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M7.5 5A2.508 2.508 0 0 0 5 7.5v33C5 41.875 6.125 43 7.5 43h33c1.375 0 2.5-1.125 2.5-2.5v-33C43 6.125 41.875 5 40.5 5h-33Zm0 1h33c.834 0 1.5.666 1.5 1.5v33c0 .834-.666 1.5-1.5 1.5h-33c-.834 0-1.5-.666-1.5-1.5v-33C6 6.666 6.666 6 7.5 6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.752 19.893v8.213h2.264m-9.627 0v-8.213h2.67c1.54 0 2.771 1.232 2.771 2.772s-1.232 2.772-2.772 2.772h-2.67m11.377 2.669l2.772-8.213l2.772 8.213m5.056-8.213L27.696 24l-2.67-4.107m2.67 8.213V24m11.915 0a4.175 4.175 0 0 1-4.175 4.175h0A4.175 4.175 0 0 1 31.26 24h0a4.175 4.175 0 0 1 4.175-4.175h0A4.175 4.175 0 0 1 39.61 24h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.981 20.16c.516 2.545-1.062 4.389-2.414 5.055m5.354 2.659c-.516-2.546 1.062-4.39 2.415-5.056"/>`),
		g.Group(children),
	)
}