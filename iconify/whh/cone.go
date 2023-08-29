package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-27 0-45.5-19T0 959.5t18.5-45T64 896h64l49-128h670l49 128h64q26 0 45 18.5t19 45t-19 45.5t-45 19zM423 128l25-64q22-64 64-64q45 0 64 64l25 64H423zm276 256H325l49-128h276zM276 512h472l50 128H226z"/>`),
		g.Group(children),
	)
}