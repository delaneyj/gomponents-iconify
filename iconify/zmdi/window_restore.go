package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 107h85V21h256v256h-85v86H0V107zm256 0v128h43V64H128v43h128zM43 192v128h170V192H43z"/>`),
		g.Group(children),
	)
}