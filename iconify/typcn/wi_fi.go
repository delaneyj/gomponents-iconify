package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.414 19.412a1.994 1.994 0 0 0 0-2.826a1.994 1.994 0 0 0-2.828-.002a2 2 0 1 0 2.828 2.828zm7.071-7.897a1.99 1.99 0 0 1-1.414-.586c-3.899-3.899-10.243-3.898-14.143 0A2 2 0 0 1 2.099 8.1c5.459-5.458 14.341-5.458 19.799 0a2 2 0 0 1-1.413 3.415zM7.757 15.757a2 2 0 0 1-1.414-3.414c3.118-3.119 8.194-3.119 11.313 0a2 2 0 0 1-2.829 2.829a4.005 4.005 0 0 0-5.657 0a1.99 1.99 0 0 1-1.413.585z"/>`),
		g.Group(children),
	)
}