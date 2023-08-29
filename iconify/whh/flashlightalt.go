package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlightalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1005 885l-120 120q-19 19-46 19t-46-19L365 576H192q-18 0-74-48T15 433Q0 418 0 397t15-36L361 15q15-15 36-15t36 15q51 51 97 105.5t46 71.5v173l429 428q19 19 19 46t-19 46zM407 104q-9-9-21.5-9t-21.5 9L104 364q-9 9-9 21.5t9 21.5t22 9t22-9l259-259q9-9 9-22t-9-22zm512 708L564 457q-9-9-21.5-9t-21.5 9t-9 22t9 21l355 355q9 9 21.5 9t21.5-9t9-21.5t-9-21.5zM780 503l-67-67q-9-8-9-21t9-22t22-9t22 9l66 67q9 9 9 21.5t-9 21.5t-21.5 9t-21.5-9z"/>`),
		g.Group(children),
	)
}