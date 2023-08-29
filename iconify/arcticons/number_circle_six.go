package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M23.562 11.85c-4.407.23-7.712 4.186-7.712 8.6V28A8.15 8.15 0 0 0 24 36.15h0A8.15 8.15 0 0 0 32.15 28"/><path d="M32.15 28A8.15 8.15 0 0 0 24 19.852h0a8.15 8.15 0 0 0-8.15 8.15m14.868-13.49c-1.49-1.94-3.358-2.662-5.958-2.662h-1.198"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}