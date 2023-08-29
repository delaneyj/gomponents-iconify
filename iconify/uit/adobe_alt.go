package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.917 2.224A.5.5 0 0 0 9.5 2h-8a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .96.194l8-19a.504.504 0 0 0-.043-.47zM2 19.024V3h6.747L2 19.024zM22.5 2h-8a.5.5 0 0 0-.46.694l8 19A.5.5 0 0 0 23 21.5v-19a.5.5 0 0 0-.5-.5zM22 19.024L15.253 3H22v16.024zm-9.532-9.7A.498.498 0 0 0 12.003 9H12a.5.5 0 0 0-.466.318l-3.5 9A.5.5 0 0 0 8.5 19h3.191l1.362 2.724A.5.5 0 0 0 13.5 22h3a.5.5 0 0 0 .468-.676l-4.5-12zM13.808 21l-1.36-2.724A.501.501 0 0 0 12 18H9.23l2.761-7.099L15.778 21h-1.97z"/>`),
		g.Group(children),
	)
}