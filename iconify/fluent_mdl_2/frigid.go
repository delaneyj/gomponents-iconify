package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frigid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1419q29 10 52 28t41 42t26 52t9 59q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75q0-30 9-58t26-53t40-42t53-28v-395h128v395zm256-100q57 63 88 141t32 163q0 90-35 167t-96 135t-141 90t-168 33q-88 0-168-33t-140-90t-96-134t-36-168q0-85 31-163t89-141V320q0-66 25-124t68-101t102-69T960 0q66 0 124 25t101 69t69 102t26 124v999zm-320 601q62 0 118-23t100-63t68-95t26-118q0-75-33-137t-87-113V320q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75v1051q-54 50-87 112t-33 138q0 63 25 118t69 95t99 63t119 23v64v-64z"/>`),
		g.Group(children),
	)
}