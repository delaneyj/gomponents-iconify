package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brokenheart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m805.694 724l-293 300l-294-300q-119-122-168.5-231T.694 256q0-106 75-181t181-75q64 0 119.5 29.5t90.5 80.5l46 274l-128 64l128 384V480l128-64l-90-293q34-57 92-90t126-33q106 0 181 75t75 181q0 127-49.5 236.5T805.695 724z"/>`),
		g.Group(children),
	)
}