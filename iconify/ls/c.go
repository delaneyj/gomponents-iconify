package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func C(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m507 242l-56 39c-38-51-98-84-167-84c-70 0-134 34-174 85c-26 34-41 75-41 119s15 85 41 119c40 51 104 85 174 85c69 0 129-33 167-84l56 39c-50 68-131 111-222 111c-92 0-178-44-230-112C20 514 0 459 0 401s20-113 55-158c52-68 138-112 230-112c91 0 172 43 222 111z"/>`),
		g.Group(children),
	)
}