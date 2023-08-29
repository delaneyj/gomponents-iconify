package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.694 1024q0-20-17.5-43t-47-47t-68-52t-81-62.5t-85-74.5t-81-91.5t-68-109t-47-132T.694 256q0-106 75-181t181-75t181 75t75 181q0-106 75-181t181-75t181 75t75 181q0 82-17.5 156.5t-47 132t-68 109t-81 91.5t-85 74.5t-81 62.5t-68 52t-47 47t-17.5 43z"/>`),
		g.Group(children),
	)
}