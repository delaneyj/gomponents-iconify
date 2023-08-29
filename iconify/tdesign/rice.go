package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4C8.855 4 6.17 5.211 5.493 7.714A14.927 14.927 0 0 0 5.057 10h13.886a14.923 14.923 0 0 0-.435-2.282C17.831 5.212 15.148 4 12 4Zm8 8H4a8 8 0 1 0 16 0ZM3.048 10c.08-.886.257-1.859.514-2.809C4.605 3.341 8.564 2 12 2c3.44 0 7.397 1.343 8.439 5.196c.256.948.433 1.92.513 2.804H22v2c0 5.523-4.477 10-10 10S2 17.523 2 12v-2h1.048Zm7.95-5.002h2.004v2.004h-2.004V4.998Zm-3 2h2.004v2.004H7.998V6.998Zm6 0h2.004v2.004h-2.004V6.998Z"/>`),
		g.Group(children),
	)
}