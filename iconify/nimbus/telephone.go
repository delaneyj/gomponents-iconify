package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telephone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.91 2.75h3a11.67 11.67 0 0 1-3.8 7.46a11.71 11.71 0 0 1-7.88 3h-.77V10.4L4 9.52l.84 1.42l.59 1.06l1.07-.62a14.36 14.36 0 0 0 4.77-4.5L12 5.73l-1.22-.64l-1-.52l1.09-1.82m0-1.25a1.26 1.26 0 0 0-1.07.61L8.75 3.93a1.24 1.24 0 0 0 .49 1.75l1 .52a13.06 13.06 0 0 1-4.36 4.1L5 8.89a1.25 1.25 0 0 0-1-.62a1.22 1.22 0 0 0-.41.07L1 9.22a1.25 1.25 0 0 0-.79 1.18v2.85a1.25 1.25 0 0 0 1.25 1.25h.75A13 13 0 0 0 15.14 2.88a1.26 1.26 0 0 0-1.25-1.38z"/>`),
		g.Group(children),
	)
}