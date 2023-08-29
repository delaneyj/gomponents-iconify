package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GithubAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 151q0 38-10.5 65T343 257.5t-40 21t-47 9.5q19 16 19 51v90H136v-65q-16 3-29.5 3t-23-2.5t-17-6.5t-12-8.5T47 341t-4-7l-1-3q-6-14-13.5-24T16 294l-5-3q-11-9-11-12.5t7-4.5h6q12 1 23 8t15 14l5 6q27 47 81 23q3-24 18-37q-27-3-47-9.5t-39.5-21T38 216t-11-65q0-43 29-74q-13-33 3-74q3 1 8-.5t25 6T136 32q33-9 70-10q36 1 70 10q23-16 42.5-23T345 2l7 1q17 41 3 74q29 31 29 74zM32 286.5q1-2.5-2.5-4t-4.5 1t2.5 4t4.5-1zM43.5 299q2.5-2-1-5.5t-6-1.5t1 5.5t6 1.5zM54 315q3-2 0-6.5t-6-2.5t0 6.5t6 2.5zm15.5 16q2.5-3-1.5-7.5t-7-1.5t1.5 7.5t7 1.5zm20.5 8.5q1-3.5-4.5-5.5t-6.5 2t4.5 5.5t6.5-2zm17 5.5q6 0 6-4t-6-4t-6 4t6 4zm22-2q3-1 4.5-2.5t.5-2.5q0-5-6-4q-3 1-4.5 2.5t-.5 3.5q0 4 6 3z"/>`),
		g.Group(children),
	)
}