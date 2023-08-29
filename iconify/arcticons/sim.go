package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M24.05 4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a1.75 1.75 0 0 0-1.2.5L8.93 18.87a1.8 1.8 0 0 0-.51 1.26h0v21.4a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V20.39a1.82 1.82 0 0 0 0-.57V6.45a2 2 0 0 0-2-1.95H24Zm-8.16 18.28h4.06m3.87 0h8.42M15.84 35.91h4.06m3.87 0h8.42m0-6.57h-4.06m-3.88 0h-8.41"/>`),
		g.Group(children),
	)
}