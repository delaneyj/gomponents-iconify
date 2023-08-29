package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floppy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m299 0l85 85v256q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h256zM192 341q27 0 45.5-18.5t18.5-45t-18.5-45.5t-45.5-19t-45.5 19t-18.5 45.5t18.5 45T192 341zm64-213V43H43v85h213z"/>`),
		g.Group(children),
	)
}