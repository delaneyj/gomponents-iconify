package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm918 1493q12 0 21-12.5t13-25t11-37.5l269-1044q7-27-9-51t-36-24h-179q-12 0-21 12.5t-13 25t-11 37.5l-134 521l-134-521q-6-24-11-37.5t-14-25.5t-20-12H806q-16 0-25.5 19.5T761 374L627 895L493 374q-7-25-11-37.5t-13-25t-21-12.5H269q-20 0-36 24t-9 51l269 1044q7 25 11 37.5t13 25t21 12.5h178q27 0 45-75l135-523l135 523q18 75 45 75h178z"/>`),
		g.Group(children),
	)
}