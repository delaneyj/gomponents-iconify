package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Periscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.102 21c1.406 0 6.985-6.329 6.985-10.571C19.087 6.368 15.915 3 12.102 3c-4.017 0-7.188 3.366-7.188 7.429C4.913 14.67 10.492 21 12.102 21zM10.979 5.885a4.696 4.696 0 0 1 1.143-.139c2.25 0 4.186 1.913 4.186 4.398c0 2.238-1.936 4.151-4.196 4.151c-2.509 0-4.444-1.913-4.444-4.151c0-1.047.338-1.98.922-2.723v.022c0 .934.755 1.676 1.688 1.676c.933.002 1.722-.764 1.722-1.697a1.68 1.68 0 0 0-1.02-1.54l-.001.003z"/>`),
		g.Group(children),
	)
}