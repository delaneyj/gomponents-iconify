package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ducky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 192q19-2 35.5-2t42.5 1t50 1q0 16-6.5 29.5T998 244t-24.5 15t-31 11t-28 6.5T888 281q-7 26-31.5 82.5T832 448q0 17 13 38t32 43.5t38 47t32 58t13 69.5q0 94-58.5 167.5T741 984.5T512 1024q-78 0-151.5-22.5t-139-70t-115-114.5t-78-162T0 448q0 43 39.5 72.5t103 42.5T288 576q47 0 116-21.5t120.5-52T576 448q0-21-20-48.5T512 348t-44-58t-20-66q0-93 65.5-158.5T672 0q64 0 117.5 34t81.5 89q38-8 82-25.5t71-33.5q0 15-7 30t-16 25.5t-25.5 23T949 161t-29 17t-24 14zm-160-64q-13 0-22.5 19t-9.5 45.5t9.5 45T736 256t22.5-18.5t9.5-45t-9.5-45.5t-22.5-19z"/>`),
		g.Group(children),
	)
}