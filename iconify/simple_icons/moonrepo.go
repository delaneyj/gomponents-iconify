package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonrepo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.514 7.342c.197.52.3 1.071.302 1.627c0 3.525-3.818 5.728-6.87 3.965a4.577 4.577 0 0 1-2.289-3.965c0-3.323 3.428-5.538 6.458-4.176a10.587 10.587 0 0 0-6.46-2.181C4.772 2.614.005 7.381 0 13.265c.002 5.883 4.772 10.653 10.655 10.655c5.883-.004 10.651-4.773 10.655-10.655c.002-2.11-.623-4.17-1.796-5.923Zm-4.673-2.676c0 3.525 3.816 5.729 6.868 3.966A4.578 4.578 0 0 0 24 4.666c0-3.526-3.816-5.727-6.87-3.967a4.585 4.585 0 0 0-2.289 3.967"/>`),
		g.Group(children),
	)
}