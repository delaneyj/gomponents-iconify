package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.141 41.537l-30.498.963l.271-28.936l30.227-1.466v29.439z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.141 12.098L32.215 5.5L6.269 7.296l5.645 6.268"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.643 42.5l-5.784-8.137l.41-27.067m14.792 29.402V19.954l12.546 16.385V19.057M18.745 36.855l4.633-.314m-4.633-16.587h2.316m10.23-.74l4.632-.315"/>`),
		g.Group(children),
	)
}