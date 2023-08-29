package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashNotAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l7.569 7.569M23.5 23.5l-2.882-2.882M12.05 2.05a2.801 2.801 0 0 0 3.136.571A2.801 2.801 0 0 0 17 0a2.8 2.8 0 0 0 1.814 2.621a2.801 2.801 0 0 0 3.136-.57a2.801 2.801 0 0 0-.571 3.135A2.801 2.801 0 0 0 24 7m-6 14.5H.5v-13H5m3.069-.431L8.159 8A2.997 2.997 0 0 0 9.5 5.5h6a3 3 0 0 0 3 3h3v10c0 .768-.289 1.47-.764 2l-.118.118M8.068 8.068l2.786 2.786m9.764 9.764l-4.471-4.471m0 0a4 4 0 0 0-5.293-5.293m5.293 5.293l-5.293-5.293M9 12.562A4 4 0 0 0 14.438 18"/>`),
		g.Group(children),
	)
}