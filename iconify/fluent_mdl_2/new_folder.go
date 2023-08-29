package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0q27 0 50 10t40 27t28 41t10 50v992q0 31 9 54t24 44t31 41t31 45t23 58t10 78v480q0 27-10 50t-27 40t-41 28t-50 10H768v-128h896v-480q0-45 9-77t24-58t31-46t31-40t23-44t10-55V128H512v768H384V0h1408zm128 1440q0-24-4-42t-13-33t-20-29t-27-32q-15 17-26 31t-20 30t-13 33t-5 42v480h128v-480zM896 1664H512v384H384v-384H0v-128h384v-384h128v384h384v128z"/>`),
		g.Group(children),
	)
}