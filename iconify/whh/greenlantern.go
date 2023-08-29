package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenlantern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M590 128q107 49 174.5 156T832 512t-67.5 228T590 896h242l64 128H0l64-128h242q-107-49-174.5-156T64 512t67.5-228T306 128H64L0 0h896l-64 128H590zM192 512q0 106 75 181t181 75t181-75t75-181t-75-181t-181-75t-181 75t-75 181z"/>`),
		g.Group(children),
	)
}