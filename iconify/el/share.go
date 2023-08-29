package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v1200h1200V573.486l-196.875 208.739v220.898h-806.25v-806.25h396.68V0H0zm857.861 0v225.977c-205.254.005-579.542 2.254-579.542 641.895c42.436-427.736 237.375-430.415 579.542-430.42v246.776L1200 342.09L857.861 0z"/>`),
		g.Group(children),
	)
}