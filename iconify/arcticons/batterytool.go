package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batterytool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.462 21.83c-.253-.01-.5-.04-.755-.04c-10.098 0-18.308 8.422-18.707 18.963c-.4-10.54-8.61-18.963-18.707-18.963c-.255 0-.502.03-.755.04c-.01.266-.038.527-.038.796c0 10.912 8.392 19.758 18.745 19.758c.255 0 .503-.03.755-.04c.253.01.5.04.755.04c10.353 0 18.745-8.846 18.745-19.758c0-.268-.028-.53-.038-.796Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 25.532c1.3-2.371 5.645-5.88 5.645-5.88A28.111 28.111 0 0 0 24 5.615a28.11 28.11 0 0 0-5.645 14.037S22.7 23.16 24 25.532Z"/>`),
		g.Group(children),
	)
}