package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderFollowing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m21 11.889l-4.455 4.455l-2.094-2.176L9.538 19H5c-1.1 0-2-.9-2-2V3h18v8.889zM16.508 22L24 14.482l-1.388-1.376l-6.094 6.094l-2.106-2.188L13 18.4l3.508 3.6zM9 14H5v-1h4v1zm-4-2h7v-1H5v1zm7-2H5V9h7v1zM5 7h14V5H5v2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}