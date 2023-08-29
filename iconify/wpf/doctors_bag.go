package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorsBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M10 1C8.355 1 7 2.355 7 4v2h2V4c0-.563.437-1 1-1h6c.563 0 1 .437 1 1v2h2V4c0-1.645-1.355-3-3-3h-6zM3 7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h20c1.656 0 3-1.344 3-3V10a3 3 0 0 0-3-3H3zm10 2.906A6.089 6.089 0 0 1 19.094 16A6.089 6.089 0 0 1 13 22.094A6.089 6.089 0 0 1 6.906 16A6.089 6.089 0 0 1 13 9.906zM12 13v2h-2v2h2v2h2v-2h2v-2h-2v-2h-2z"/>`),
		g.Group(children),
	)
}