package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M145 216q0-68 68.5-68t68.5 68t-68.5 68t-68.5-68zm4-213h128l39 42h68q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67zm64 320q44 0 75.5-31.5T320 216t-31.5-75.5T213 109t-75 31.5t-31 75.5t31 75.5t75 31.5z"/>`),
		g.Group(children),
	)
}