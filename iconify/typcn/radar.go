package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20c3.86 0 7-3.141 7-7s-3.14-7-7.003-7C8.139 6 5 9.141 5 13s3.14 7 7 7zM11 8.102V10a1 1 0 1 0 2 0V8.102A5.013 5.013 0 0 1 16.899 12H15a1 1 0 1 0 0 2h1.899A5.013 5.013 0 0 1 13 17.898V16a1 1 0 1 0-2 0v1.898A5.013 5.013 0 0 1 7.101 14H9a1 1 0 1 0 0-2H7.101A5.012 5.012 0 0 1 11 8.102z"/>`),
		g.Group(children),
	)
}