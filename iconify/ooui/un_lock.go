package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 8V5s0-5-5-5a4.63 4.63 0 0 0-4.88 4h2C7.31 2.93 8 2 10 2c3 0 3 2 3 3.5V8H3.93A1.93 1.93 0 0 0 2 9.93v8.15A1.93 1.93 0 0 0 3.93 20h12.14A1.93 1.93 0 0 0 18 18.07V9.93A1.93 1.93 0 0 0 16.07 8zm-5 8a2 2 0 1 1 2-2a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}