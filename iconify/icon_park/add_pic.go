package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddPic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4"><path d="M38 21V40C38 41.1046 37.1046 42 36 42H8C6.89543 42 6 41.1046 6 40V12C6 10.8954 6.89543 10 8 10H26.3636"/><path fill="#2F88FF" d="M12.0005 31.0308L18.0005 23L21.0005 26L24.5005 20.5L32.0005 31.0308H12.0005Z"/><path d="M34.0005 10H42.0005"/><path d="M37.9946 5.79541V13.7954"/></g>`),
		g.Group(children),
	)
}