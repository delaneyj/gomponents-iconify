package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1451 43q9 0 44 29t82 73t100 97t97 100t74 82t29 45q0 7-17 29t-44 53t-64 70t-74 76t-76 74t-69 63t-53 45t-29 17q-20 0-33-14t-14-35q0-3 4-21t10-43t13-55t15-56t12-47t7-28H725q-26 0-49 10t-40 28t-28 41t-11 49v1323H341V725q0-79 30-148t83-122t122-83t149-31h740l-7-28q-5-20-12-46t-14-56t-14-55t-10-44t-4-22q0-20 13-33t34-14z"/>`),
		g.Group(children),
	)
}