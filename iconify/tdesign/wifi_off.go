package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21L21 22.414l-7.304-7.304l-2.554-2.554a6.467 6.467 0 0 0-3.739 1.847l-.707.707l-1.414-1.414l.707-.707a8.466 8.466 0 0 1 3.48-2.106l-.37-.37l-3.067-3.067c-1.16.6-2.25 1.386-3.225 2.36l-.707.708L.686 9.1l.707-.707A15.016 15.016 0 0 1 4.557 5.97L1.586 3L3 1.586Zm18.192 8.222a12.978 12.978 0 0 0-10.934-3.692l-.99.134L9 4.267l.991-.133a14.978 14.978 0 0 1 12.614 4.26l.707.706l-1.414 1.415l-.707-.707Zm-10.606 7.778a2 2 0 0 1 2.828 0l.707.707l-2.12 2.121l-2.122-2.121l.707-.707Z"/>`),
		g.Group(children),
	)
}