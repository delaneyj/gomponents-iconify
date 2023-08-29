package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336zm-685 1089q14 19 24 30.5t35.5 24.5t57.5 13h149q31 0 52.5-21.5t21.5-52.5V373q0-31-21.5-52.5T1373 299h-149q-31 0-53 21.5t-22 52.5v671L685 367q-13-19-24-31t-36-24.5t-57-12.5H419q-31 0-52.5 21.5T345 373v1046q0 31 21.5 52.5T419 1493h149q31 0 53-21.5t22-52.5V748z"/>`),
		g.Group(children),
	)
}