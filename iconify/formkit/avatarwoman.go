package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avatarwoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 10h5c1.66 0 3 1.34 3 3v2H1v-2c0-1.66 1.34-3 3-3Zm0-6h5v2.5a2.5 2.5 0 0 1-5 0V4Z"/><path fill="currentColor" d="M9.5 3.93V3.5c0-1.1-.9-2-2-2L7 1H5.5c-1.1 0-2 .9-2 2v.93c0 1.33-.53 2.6-1.46 3.54l-.54.54l.74.37c.76.38 1.67.23 2.26-.37h4c.6.6 1.51.75 2.26.37l.74-.37l-.54-.54A4.994 4.994 0 0 1 9.5 3.93Z"/><path fill="currentColor" d="M7 1s.5 0 .5.5H7V1Z"/>`),
		g.Group(children),
	)
}