package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 50h-66V16a6 6 0 0 0-12 0v34H56a30 30 0 0 0-30 30v112a30 30 0 0 0 30 30h144a30 30 0 0 0 30-30V80a30 30 0 0 0-30-30Zm18 142a18 18 0 0 1-18 18H56a18 18 0 0 1-18-18V80a18 18 0 0 1 18-18h144a18 18 0 0 1 18 18Zm-54-54H92a26 26 0 0 0 0 52h72a26 26 0 0 0 0-52Zm-26 12v28h-20v-28Zm-60 14a14 14 0 0 1 14-14h14v28H92a14 14 0 0 1-14-14Zm86 14h-14v-28h14a14 14 0 0 1 0 28Zm-90-70a10 10 0 1 1 10 10a10 10 0 0 1-10-10Zm88 0a10 10 0 1 1 10 10a10 10 0 0 1-10-10Z"/>`),
		g.Group(children),
	)
}