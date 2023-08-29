package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.604 6.193a6.519 6.519 0 1 1 9.509 8.913l-9.58 9.672a.75.75 0 0 1-1.066 0l-9.58-9.672a6.518 6.518 0 0 1-.263-8.892c2.588-2.943 7.17-2.953 9.772-.021l.604.68l.604-.68Zm8.646 1.011a5.02 5.02 0 0 0-7.524-.016L14.56 8.501a.75.75 0 0 1-1.122 0l-1.165-1.313a5.02 5.02 0 1 0-7.321 6.863L14 23.184l9.047-9.133a5.019 5.019 0 0 0 .203-6.847Z"/>`),
		g.Group(children),
	)
}