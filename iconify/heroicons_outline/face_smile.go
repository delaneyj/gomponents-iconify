package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.182 15.182a4.5 4.5 0 0 1-6.364 0M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0ZM9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75S9.168 9 9.375 9s.375.336.375.75Zm-.375 0h.008v.015h-.008V9.75Zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75s.168-.75.375-.75s.375.336.375.75Zm-.375 0h.008v.015h-.008V9.75Z"/>`),
		g.Group(children),
	)
}