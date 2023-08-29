package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openmined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0a1.43 1.43 0 0 0-1.25.725L7.082 7.082L.725 10.75a1.44 1.44 0 0 0 0 2.5l6.357 3.668l3.668 6.357a1.44 1.44 0 0 0 2.5 0l3.668-6.357l6.357-3.668c.967-.544.967-1.945 0-2.5l-6.357-3.668L13.25.725A1.429 1.429 0 0 0 12 0zm.006 4.237l7.757 7.769l-7.769 7.757l-7.757-7.757zm-.012 3.112l-4.656 4.657l4.656 4.656l4.657-4.656z"/>`),
		g.Group(children),
	)
}