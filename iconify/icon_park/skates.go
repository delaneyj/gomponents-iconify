package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 5H31V13H5V5Z"/><path fill="#2F88FF" d="M9 36V13H27V23C41 23 41 32 41 36H9Z"/><rect width="36" height="6" x="7" y="36"/></g>`),
		g.Group(children),
	)
}