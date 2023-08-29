package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SippyCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSippyCup0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m34 44l3-28H11l3.5 28H34ZM24 10l-2-6M6 16h36m-5.005 9s5.47 0 4.973 4.404C41.471 33.81 36 32.928 36 32.928M11.005 25s-5.47 0-4.973 4.404C6.529 33.81 12 32.928 12 32.928"/><path fill="#fff" d="M37 10H11v6h26v-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSippyCup0)"/>`),
		g.Group(children),
	)
}