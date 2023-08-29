package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1zM400 1208q-7-19-9.5-30t-1-24.5t14.5-20t38-6.5h143q15 0 21 1.5t15.5 11.5t24.5 34q32 53 98.5 85t147.5 32q111 0 190-57.5t79-139.5q0-81-79-138.5T892 898q-95 0-168 43q-9 5-16.5 10.5T697 959t-10 3t-21 1H507q-38 0-66-35.5T413 847l30-479q0-29 22-49t52-20h766q29 0 50.5 21t21.5 50v77q0 29-21.5 50t-50.5 21H683l-14 217q107-40 223-40q139 0 256.5 53.5t186 145.5t68.5 200q0 109-68.5 201t-186 145t-256.5 53q-97 0-194.5-30.5t-181-97.5T400 1208z"/>`),
		g.Group(children),
	)
}