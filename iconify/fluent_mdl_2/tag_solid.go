package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0h896v896L896 1920L0 1024L1024 0zm448 624q36 0 68-14t56-38t38-56t14-68q0-36-14-68t-38-56t-56-38t-68-14q-36 0-68 14t-56 38t-38 56t-14 68q0 36 14 68t38 56t56 38t68 14z"/>`),
		g.Group(children),
	)
}