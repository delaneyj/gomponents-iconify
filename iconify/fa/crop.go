package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M557 1152h595V557zm-45-45l595-595H512v595zm1152 77v192q0 14-9 23t-23 9h-224v224q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23v-224H288q-14 0-23-9t-9-23V512H32q-14 0-23-9t-9-23V288q0-14 9-23t23-9h224V32q0-14 9-23t23-9h192q14 0 23 9t9 23v224h851L1609 9q10-9 23-9t23 9q9 10 9 23t-9 23l-247 246v851h224q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}