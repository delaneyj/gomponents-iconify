package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalanceWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 320v21q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v21H192q-18 0-30.5 12.5T149 107v170q0 18 12.5 30.5T192 320h192zm-192-43V107h213v170H192zm85.5-53q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T245 192t9.5 22.5t23 9.5z"/>`),
		g.Group(children),
	)
}