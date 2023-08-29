package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anymemo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.1 26.3L24 7.1L12.9 26.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.3 26.3v-3.5l1.6-2.4h2.2l1.6 2.4v3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.8 36.4l2.6 4.5h10.1l-8.4-14.6h-8.4l-1.6 2.5h-2.2l-1.6-2.5h-8.4L4.5 40.9h10.1l2.6-4.5l2.5 4.5h8.6l2.5-4.5z"/>`),
		g.Group(children),
	)
}