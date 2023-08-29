package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func D(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 12a2 2 0 0 1 2-2h8c7.732 0 14 6.268 14 14s-6.268 14-14 14h-8a2 2 0 0 1-2-2V12Zm4 2v20h6c5.523 0 10-4.477 10-10s-4.477-10-10-10h-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}