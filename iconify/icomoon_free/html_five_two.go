package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFiveTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.946 0L2.23 14.4L7.992 16l5.777-1.602L15.055 0H.947zm11.722 13.482l-4.644 1.287v.007l-.012-.004l-.012.004v-.007l-4.644-1.287L2.258 1.178h11.508l-1.098 12.304zm-2.5-5.198l-.204 2.29l-1.972.532l-1.967-.53l-.126-1.41H4.126l.247 2.774l3.626 1.003l3.615-1.003l.485-5.422H5.662l-.161-1.809h6.758l.158-1.766H3.57l.477 5.341z"/>`),
		g.Group(children),
	)
}