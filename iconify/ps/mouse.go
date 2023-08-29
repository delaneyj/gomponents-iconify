package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 0Q106 0 55.5 50T5 171v170q0 71 50.5 121T176 512t120.5-50T347 341V171q0-71-50.5-121T176 0zm128 341q0 53-37.5 90.5T176 469t-90.5-37.5T48 341v-85h256v85zm0-128H48v-42q0-53 37.5-90.5T176 43t90.5 37.5T304 171v42zM176 85q-21 0-21 22v64q0 21 21 21t21-21v-64q0-22-21-22z"/>`),
		g.Group(children),
	)
}