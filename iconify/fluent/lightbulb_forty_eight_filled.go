package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 4C15.992 4 9.5 10.492 9.5 18.5c0 4.235 1.817 8.047 4.71 10.696c.448.41.748.873.87 1.349l1.27 4.955h15.3l1.27-4.955c.122-.476.422-.938.87-1.35A14.465 14.465 0 0 0 38.5 18.5C38.5 10.492 32.008 4 24 4Zm7.008 34H16.992l.623 2.43a4.75 4.75 0 0 0 4.6 3.57h3.57a4.75 4.75 0 0 0 4.6-3.57l.623-2.43Z"/>`),
		g.Group(children),
	)
}