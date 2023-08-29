package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 99V56q0-11-11-11h-42q-11 0-11 11v43q0 10 11 10h42q11 0 11-10zM53 387h320q11 0 11-11V195h-45q2 12 2 21q0 53-37.5 90.5T213 344t-90.5-37.5T85 216q0-11 2-21H43v181q0 11 10 11zm160.5-256q-35.5 0-60.5 25t-25 60t25 60t60.5 25t60.5-25t25-60t-25-60t-60.5-25zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h341z"/>`),
		g.Group(children),
	)
}