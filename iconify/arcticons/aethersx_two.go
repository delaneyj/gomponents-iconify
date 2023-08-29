package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AethersxTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.672 4.477L5.907 11.615a.367.367 0 0 0-.214.372l2.108 20.476a.492.492 0 0 0 .262.386l18.14 9.462a.374.374 0 0 0 .429-.058l12.722-11.87a.707.707 0 0 0 .22-.443l2.052-19.709a.315.315 0 0 0-.227-.335L22.168 4.44a.724.724 0 0 0-.496.037Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.798 11.902l21.005 7.667l-.308 22.76m14.993-32.113l-14.644 9.51"/>`),
		g.Group(children),
	)
}