package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22a9 9 0 1 1 0-18a9 9 0 0 1 0 18Zm0-2a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm1-7h3v2h-5V8h2v5ZM1.745 6.283l3.536-3.536l1.414 1.414L3.16 7.697L1.746 6.283Zm16.97-3.536l3.536 3.536l-1.414 1.414l-3.536-3.536l1.415-1.414Z"/>`),
		g.Group(children),
	)
}