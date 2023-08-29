package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Req(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#00E6A0"/><path fill="#FFF" d="M23.073 7.5a.653.653 0 0 1 .477.192l-.006-.006a.65.65 0 0 1 .184.364l.009.104l-.007 3.417c0 .449-.178.88-.496 1.197l-4.24 4.227l4.537 4.523a1.689 1.689 0 1 1-2.383 2.394l-5.735-5.717a1.686 1.686 0 0 1 0-2.395l5.022-5.012h-8.498v11.889a1.69 1.69 0 0 1-1.718 1.718l-.17-.006A1.696 1.696 0 0 1 8.5 22.677V9.347c0-1.066.763-1.847 1.796-1.847z"/></g>`),
		g.Group(children),
	)
}