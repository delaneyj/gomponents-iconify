package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-384-896q-95 0-159.5 56t-64.5 136t61 136t147 56q5 0 16-1v1q0 18 3.5 31t8 19t9 9.5t8.5 4.5h3h-64q-48 0-100 21t-88 59t-36 80q0 77 59.5 118.5t196.5 41.5q100 0 146-44t46-116q0-29-13-56.5t-32-47t-38-36.5t-32-30.5t-13-21.5q0-28 7-50q54-23 87.5-70t33.5-104q0-74-47-128h47l64-64h-256zm0 320q-44 0-78.5-36.5t-46.5-91.5q-11-52 18-90t75-38q44 0 78.5 36.5t46.5 91.5q11 52-18 90t-75 38zm128 288q0 40-42 68t-102 28t-102-28t-42-68q0-39 59-67.5t133-28.5q28 0 62 29.5t34 66.5z"/>`),
		g.Group(children),
	)
}