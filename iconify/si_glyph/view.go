package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func View(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 4)"><path d="M8 0C3.598 0 .031 2.66.031 3.969C.031 5.278 3.597 7.938 8 7.938c4.4 0 7.969-2.618 7.969-3.969S12.4 0 8 0zm-.01 7.062C4.342 7.062 2.869 5.011 2.869 4C2.869 2.989 4.342.938 7.99.938c3.646 0 5.119 2.02 5.119 3.062c0 1.042-1.472 3.062-5.119 3.062z"/><ellipse cx="7.932" cy="3.963" rx="1.932" ry="1.963"/></g>`),
		g.Group(children),
	)
}