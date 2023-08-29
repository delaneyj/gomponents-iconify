package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 768q40 0 75 15t61 41t41 61t15 75q0 40-15 75t-41 61t-61 41t-75 15h-544l-384 768H662l256-768H256l-67 128H0l108-320L0 640h189l67 128h662L662 0h266l384 768h544z"/>`),
		g.Group(children),
	)
}