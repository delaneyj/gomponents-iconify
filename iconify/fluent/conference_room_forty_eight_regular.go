package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConferenceRoomFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m22.503 7.776l14.5 3c.58.12.997.632.997 1.224v24a1.25 1.25 0 0 1-.997 1.224l-14.5 3A1.25 1.25 0 0 1 21 39.002V9a1.25 1.25 0 0 1 1.503-1.225ZM18 10v2.5h-5.5v23H18V38h-6.75a1.25 1.25 0 0 1-1.243-1.122L10 36.751v-25.5c0-.648.492-1.18 1.122-1.244L11.25 10H18Zm5.5.536v26.93l12-2.483V13.018l-12-2.482Zm3 11.964a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/>`),
		g.Group(children),
	)
}