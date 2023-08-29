package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BucketColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1690 960l-858 858l-730-730l666-667V192q0-40 15-75t41-61t61-41t75-15q40 0 75 15t61 41t41 61t15 75v640h-128V192q0-26-19-45t-45-19q-26 0-45 19t-19 45v283l-549 549h1098l65-64l-211-211l90-90l301 301zm-858 678l485-486H347l485 486zm1147-48q20 35 30 74t10 80q0 61-22 116t-61 97t-92 66t-116 25q-62 0-116-24t-94-64t-63-95t-24-116q0-79 40-148l257-450l251 439zm-251 330q36 0 66-14t52-39t34-57t12-67q0-49-24-90l-140-244l-146 256q-23 40-23 84q0 35 13 66t37 54t53 37t66 14z"/>`),
		g.Group(children),
	)
}