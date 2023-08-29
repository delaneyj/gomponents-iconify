package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.075 4.1c0-1.189 1.182-2.175 2.425-2.175c1.243 0 2.425.986 2.425 2.175c0 1.099-.557 1.614-1.306 2.279l-.031.027C7.845 7.065 6.925 7.88 6.925 9.5a.575.575 0 1 0 1.15 0c0-1.085.554-1.594 1.307-2.26l.02-.02c.748-.662 1.673-1.482 1.673-3.12C11.075 2.128 9.219.775 7.5.775S3.925 2.128 3.925 4.1a.575.575 0 1 0 1.15 0ZM7.5 13.358a.875.875 0 1 0 0-1.75a.875.875 0 0 0 0 1.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}