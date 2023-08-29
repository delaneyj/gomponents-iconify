package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1455 0H336Q197 0 98.5 98.5T0 336v1119q0 139 98 237.5t238 98.5h1119q139 0 237.5-98.5T1791 1455V336q0-139-98.5-237.5T1455 0zM346 1437v-11q0-70 59.5-155.5T557 1108t200.5-144T967 857q86-35 126.5-71t41.5-80q1-67-77-119t-157-52q-46 0-104 25t-107 66.5t-68 84.5q-3 6-6.5 14.5T611 736t-3.5 6t-3.5 4t-5 1t-7 0t-10.5-1H404q-30 0-48-26t-10-51q55-168 212-269t354-101q87 0 173.5 26t156.5 73t113.5 119.5T1399 674q0 81-24 145t-65.5 106.5t-92 75.5t-109 60t-112 52.5t-106.5 61t-86 76.5h595q19 0 33.5 16.5t14.5 39.5v130q0 23-14.5 39.5T1399 1493H394q-19 0-33.5-16.5T346 1437z"/>`),
		g.Group(children),
	)
}