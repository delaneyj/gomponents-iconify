package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 475v-43h170v43H43zM213 6q18 0 30.5 12T256 48v299q0 17-12.5 29.5T213 389H43q-18 0-30.5-12.5T0 347V48q0-18 12.5-30.5T43 5zm0 298V91H43v213h170z"/>`),
		g.Group(children),
	)
}