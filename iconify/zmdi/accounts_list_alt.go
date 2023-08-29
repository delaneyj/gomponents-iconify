package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountsListAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 0v43H43V0h341zM43 512v-43h341v43H43zM384 85q18 0 30.5 12.5T427 128v256q0 18-12.5 30.5T384 427H43q-18 0-30.5-12.5T0 384V128q0-18 12.5-30.5T43 85h341zm-171 59q-20 0-34 14t-14 34t14 34t34 14t34-14t14-34t-14-34t-34-14zm107 219v-32q0-24-36.5-39t-70-15t-70 15t-36.5 39v32h213z"/>`),
		g.Group(children),
	)
}