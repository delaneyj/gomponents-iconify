package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M3 18L19 2v10h42v12H19v10zm11 28.3c0-3.5-2.7-6.3-6-6.3H2v22h6c3.3 0 6-2.8 6-6.3c0-1.8-.7-3.5-1.9-4.7c1.2-1.2 1.9-2.8 1.9-4.7M8 58.6H5.1v-5.9H8c1.7 0 3 1.3 3 3s-1.3 2.9-3 2.9m0-9.3H5.1v-5.9H8c1.7 0 3 1.3 3 3s-1.3 2.9-3 2.9m32 9.3c-1.7 0-3-1.3-3-3v-9.3c0-1.6 1.3-3 3-3s3 1.3 3 3h3c0-3.5-2.7-6.3-6-6.3s-6 2.8-6 6.3v9.3c0 3.5 2.7 6.3 6 6.3s6-2.8 6-6.3h-3c0 1.7-1.3 3-3 3M62 40h-3.2L55 49.5h-2V40h-3v22h3v-9.5h2l3.8 9.5H62l-4.4-11zM27 62h3l-4-22h-5l-3 22h3l.8-6.1h4.1L27 62m-4.7-9.4l1.3-9.4l1.7 9.4h-3"/>`),
		g.Group(children),
	)
}