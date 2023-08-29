package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.831 4.059a.49.49 0 0 0-.394-.121l-11 1.25A.5.5 0 0 0 7 5.684V16c-1.654 0-3 1.122-3 2.5S5.346 21 7 21s3-1.122 3-2.5v-7.609l6-.625V14c-1.654 0-3 1.122-3 2.5s1.346 2.5 3 2.5s3-1.122 3-2.5V4.434a.501.501 0 0 0-.169-.375z"/>`),
		g.Group(children),
	)
}