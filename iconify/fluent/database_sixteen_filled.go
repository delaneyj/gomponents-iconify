package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3.5C13 4.88 10.761 6 8 6S3 4.88 3 3.5S5.239 1 8 1s5 1.12 5 2.5Zm-10 9V5.487C4.057 6.413 5.864 7 8 7s3.943-.587 5-1.513V12.5c0 1.425-2.149 2.5-5 2.5s-5-1.075-5-2.5Z"/>`),
		g.Group(children),
	)
}