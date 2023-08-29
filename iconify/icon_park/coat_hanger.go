package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoatHanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 13C19 10.2386 21.2386 8 24 8C26.7614 8 29 10.2386 29 13C29 15.7614 26.7614 18 24 18V22"/><path fill="#2F88FF" d="M44 36H4C4 31 10 22 24 22C38 22 44 31 44 36Z"/></g>`),
		g.Group(children),
	)
}