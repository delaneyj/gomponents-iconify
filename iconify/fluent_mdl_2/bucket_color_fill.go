package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BucketColorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M192 1088h1280l-640 640l-640-640zm1731 534q32 56 32 122q0 47-16 90t-47 76t-71 54t-89 20q-49 0-92-18t-75-50t-51-75t-19-92q0-62 31-116l202-353l195 342z"/>`),
		g.Group(children),
	)
}