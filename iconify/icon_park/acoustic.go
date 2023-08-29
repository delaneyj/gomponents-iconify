package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acoustic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M24 3.99976V43.9998"/><path d="M34 11.9998V35.9998"/><path d="M4 17.9998V29.9998"/><path d="M44 17.9998V29.9998"/><path d="M14 11.9998V35.9998"/></g>`),
		g.Group(children),
	)
}