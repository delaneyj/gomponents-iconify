package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M.5 7.5v27c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-27c0-2.46-.471-3-3-3h-34c-2.48 0-3 .43-3 3zm35.29 23H5.23c3.34-4.87 9.279-12.99 10.789-12.99c1.461 0 6.42 6.561 8.661 8.87c0 0 2.881-3.851 4.391-3.851c1.538 0 6.669 7.931 6.719 7.971zm-8.979-17c0-2.04 1.649-3.689 3.689-3.689s3.689 1.649 3.689 3.689s-1.649 3.689-3.689 3.689s-3.689-1.649-3.689-3.689z"/>`),
		g.Group(children),
	)
}