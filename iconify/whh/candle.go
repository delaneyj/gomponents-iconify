package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-27 0-45.5-19T0 959.5t18.5-45T64 896h896q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19zM320 832V480q0-13 9.5-22.5T352 448h320q13 0 22.5 9.5T704 480v352H320zm213-448q22-24 32.5-45t10.5-51q0-16-14.5-46T532 190l-15-22q-1 10-4.5 25.5t-18.5 48t-35 46.5q-11 7-11.5 24.5t5.5 36t16 35.5q-18 0-59-43q-26-29-26-65q0-53 43-96q22-21 27-66t-1-80l-5-34q4 2 11.5 5.5t28 16.5T526 49t39.5 36t35.5 44q20 32 30 69t9 66q0 5-.5 19.5t-1 21.5t-1.5 17t-3 16.5t-5 9.5q-13 15-44.5 25.5T533 384z"/>`),
		g.Group(children),
	)
}