package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraBw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 43q18 0 30.5 12.5T427 85v256q0 18-12.5 30.5T384 384H43q-18 0-30.5-12.5T0 341V85q0-17 12.5-29.5T43 43h68l38-43h128l39 43h68zm0 298V85H213v22q-44 0-75 31t-31 75.5t31 75.5t75 31v21h171zm-64-127.5q0 44.5-31 75.5t-76 31v-38q29 0 49-20t20-48.5t-20-48.5t-49-20v-38q45 0 76 31t31 75.5zm-175 0q0-28.5 20-48.5t48-20v137q-28 0-48-20t-20-48.5z"/>`),
		g.Group(children),
	)
}