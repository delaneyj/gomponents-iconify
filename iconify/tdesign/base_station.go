package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.344 3.93l-.707.707a9 9 0 0 0 0 12.728l.707.707l-1.414 1.414l-.707-.707c-4.296-4.296-4.296-11.26 0-15.557l.707-.707L6.344 3.93ZM19.07 2.515l.708.707c4.295 4.296 4.295 11.261 0 15.557l-.707.707l-1.415-1.414l.707-.707a9 9 0 0 0 0-12.728l-.707-.707l1.415-1.415ZM9.526 7.111l-.707.707a4.5 4.5 0 0 0 0 6.364l.707.707l-1.414 1.415l-.708-.707a6.5 6.5 0 0 1 0-9.193l.708-.707l1.414 1.414Zm6.363-1.414l.707.707a6.5 6.5 0 0 1 0 9.193l-.707.707l-1.414-1.415l.707-.707a4.5 4.5 0 0 0 0-6.364l-.707-.707l1.414-1.414ZM10 11a2 2 0 1 1 3 1.732V23h-2V12.732A2 2 0 0 1 10 11Z"/>`),
		g.Group(children),
	)
}