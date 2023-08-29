package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snacks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSnacks0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M6 14h36V8h-4l-2-4H12l-2 4H6v6Z"/><path stroke-linecap="round" d="m36 44l2-30H10l2 30h24Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSnacks0)"/>`),
		g.Group(children),
	)
}