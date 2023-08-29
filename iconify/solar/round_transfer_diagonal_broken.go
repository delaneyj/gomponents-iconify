package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundTransferDiagonalBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m6.5 4l7.378 8V7.5m3.5 12.378l-7.5-7.879v4.5"/><path d="M17 3.338A9.954 9.954 0 0 0 12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10c0-1.821-.487-3.53-1.338-5"/></g>`),
		g.Group(children),
	)
}