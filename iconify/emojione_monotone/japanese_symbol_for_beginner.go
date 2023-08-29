package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseSymbolForBeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52.117 2.118a1.49 1.49 0 0 0-1.585.269L32 19.328L13.467 2.387a1.491 1.491 0 0 0-1.584-.269A1.428 1.428 0 0 0 11 3.429v39.285c0 .396.167.772.463 1.043l19.535 17.856A1.48 1.48 0 0 0 32 62c.359 0 .72-.129 1.002-.387l19.535-17.856c.295-.271.463-.648.463-1.043V3.429a1.43 1.43 0 0 0-.883-1.311M13.93 42.096V6.726l17.068 15.603c.007.007.018.009.025.016v35.377L13.93 42.096m36.139 0L32.977 57.721V22.344c.008-.007.018-.009.025-.016L50.069 6.726v35.37"/>`),
		g.Group(children),
	)
}