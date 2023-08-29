package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkPolicy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1ZM3 7a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm10 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1ZM10.69 19.883a9.57 9.57 0 0 1-.59-2.201a5.961 5.961 0 0 1-3.995-6.777L9.6 14.4v.8a1.585 1.585 0 0 0 .4 1.045V12H8.8v-1.6h1.6a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a6.027 6.027 0 0 1 2.353 1.857L18 7.833l1.08.45a7.997 7.997 0 1 0-8.39 11.6ZM7 21a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm5 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-9-5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-2-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm23 1.5v3.863A8.169 8.169 0 0 1 18 24a8.12 8.12 0 0 1-6-7.596V12.5l6-2.5Z"/>`),
		g.Group(children),
	)
}