package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M830 526q62 34 95.5 83t36 101T951 810t-39 90q-33 54-80 86.5t-101.5 37T619 998q-13-7-21.5-20.5t-13-28t-6-39T577 869t.5-50t.5-51q-1-78-13-139q-25 11-51 11q-24 0-47-9q-64 74-106 101q-52 33-104.5 35.5t-94-21t-73-57.5T35 613Q7 560 1.5 506.5t15-100.5T84 329q16-7 31-8t38 7.5t43 19.5t46.5 27.5t50 31.5t51.5 30.5t52 25.5q10-24 29-43q-39-114-39-176q0-53 18-95.5T451.5 79t68-45.5t77-26T674 0q96 0 160 48.5T898 183q0 29-33 59t-75.5 53T688 351.5T594 413q27 22 39 53q136 26 197 60zm-316.5-78q-26.5 0-45 19T450 512.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}