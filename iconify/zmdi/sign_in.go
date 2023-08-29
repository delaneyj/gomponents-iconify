package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m151 269l55-56H0v-42h206l-55-56l30-30l107 107l-107 107zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341v-85h43v85h298V43H43v85H0V43q0-18 12.5-30.5T43 0h298z"/>`),
		g.Group(children),
	)
}