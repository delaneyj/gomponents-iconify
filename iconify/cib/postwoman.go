package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postwoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.38 12.38C3.38 20.614 10.078 32 16 32c5.573 0 12.62-11.385 12.62-19.62C28.625 4.151 22.969 0 16 0S3.38 4.151 3.38 12.38zm15.901 6.089a6.203 6.203 0 0 1 5.677-1.682a6.206 6.206 0 0 1-1.688 5.677a6.197 6.197 0 0 1-5.672 1.682a6.19 6.19 0 0 1 1.682-5.677zM7.042 16.786a6.203 6.203 0 0 1 5.677 1.682a6.187 6.187 0 0 1 1.682 5.677a6.197 6.197 0 0 1-5.672-1.682a6.206 6.206 0 0 1-1.688-5.677z"/>`),
		g.Group(children),
	)
}