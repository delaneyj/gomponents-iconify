package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DefectSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1192 42q107 37 194 102t149 152t95 191t34 217q0 144-55 267t-156 226q-71 73-113 155t-53 184H633q-11-102-52-184t-113-154q-100-103-156-226t-56-268q0-97 25-187t71-169t110-142T604 96t168-71T960 0q11 0 21 1t22 3L623 384l448 448l-263 263l113 114l376-377l-448-448l343-342zM640 1664h640v192q0 40-15 75t-41 61t-61 41t-75 15H832q-40 0-75-15t-61-41t-41-61t-15-75v-192z"/>`),
		g.Group(children),
	)
}