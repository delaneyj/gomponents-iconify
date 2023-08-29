package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TradeMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M2 2v7.529h10.313V34h7.5V9.529h10.312V2zm52.5 0l-6.562 13.177L41.375 2h-7.5v32h7.5V20.823L47.938 34L54.5 20.823V34H62V2z"/>`),
		g.Group(children),
	)
}