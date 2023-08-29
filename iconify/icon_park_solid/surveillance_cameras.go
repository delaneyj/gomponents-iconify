package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCameras(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSurveillanceCameras0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 6H5m18 17V6"/><path fill="#fff" d="m8.425 19.58l31.876 8.54l-1.743 2.64l-4.002 7.21l-1.743 2.639l-28.011-7.506L8.425 19.58Z"/><path d="m38.558 30.76l3.864 1.035l-2.07 7.727l-5.796-1.552"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSurveillanceCameras0)"/>`),
		g.Group(children),
	)
}