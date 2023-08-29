package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyHeartExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#ff5a79"><circle cx="32" cy="52" r="10"/><path d="M57.4 11.8C52.9-2.7 34.3 1.2 32 8.9C29.7 1.2 11.1-2.7 6.6 11.8C4.6 18.1 7.7 24 13 27.7c7.1 5 15.5 5 19 14.3c3.5-9.3 11.8-9.3 19-14.3c5.3-3.7 8.4-9.6 6.4-15.9"/></g>`),
		g.Group(children),
	)
}