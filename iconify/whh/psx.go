package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm246.5-758.5Q748 255 733 255t-25 10L512 461L317 266q-10-10-25-10t-25.5 10.5t-10.5 25t11 25.5l194 195l-194 194q-11 11-11 26t10.5 25.5T292 768t26-11l194-194l195 194q10 11 25 11t25.5-10.5T768 732t-11-25L563 512l195-196q11-10 11-25t-10.5-25.5z"/>`),
		g.Group(children),
	)
}