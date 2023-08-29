package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkunreadMailbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 128q18 0 30.5 12.5T427 171v256q0 17-12.5 29.5T384 469H43q-18 0-30.5-12.5T0 427V171q0-18 12.5-30.5T43 128h42V0h171v85H128v171h43V128h213z"/>`),
		g.Group(children),
	)
}