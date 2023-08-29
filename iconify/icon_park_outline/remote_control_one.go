package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M11 5.446C11 4.647 11.647 4 12.446 4h23.108C36.353 4 37 4.647 37 5.446V31c0 7.18-5.82 13-13 13s-13-5.82-13-13V5.446Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 16h3m8 0h3"/><circle cx="17" cy="10" r="2" fill="currentColor"/><circle cx="24" cy="31" r="7" stroke="currentColor" stroke-width="4"/></g>`),
		g.Group(children),
	)
}