package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WirelessChargeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11.25 22a.75.75 0 0 0 1.5 0h-1.5Zm0-1v1h1.5v-1h-1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.857 7L10 10h4l-2.857 3"/><path stroke="currentColor" stroke-width="1.5" d="M13.5 18v1.5c0 .466 0 .699-.076.883a1 1 0 0 1-.541.54C12.699 21 12.466 21 12 21s-.699 0-.883-.076a1 1 0 0 1-.54-.541c-.077-.184-.077-.417-.077-.883V18"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M4.582 7A8 8 0 1 0 9 2.582"/></g>`),
		g.Group(children),
	)
}