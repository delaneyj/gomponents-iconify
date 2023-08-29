package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleEarth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M222 121q114 108 165 114q1-11 1-19q0-25-7-50q-4 9-11 10t-15.5-5.5T339 156t-14.5-18.5t-10-15t-3.5-6.5q-47-66-163-62q-32 13-56 36q65-30 130 31zm143 182q11-20 16-39q-33-3-85.5-29.5T208 183l-35-25q-74-58-127 9q-8 24-8 49q0 38 16 73q9-26 25-26q15 0 40.5 13.5T161 295q10 3 31 10l31.5 10.5L250 322l30 3q12 0 22-1.5t20-4.5t15.5-4.5t15.5-6t12-5.5zm-152 88q76 0 128-56q-45 13-83.5 13t-62.5-7l-25-8q-26-8-31 6t7 38q32 14 67 14zm0-388q88 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213 3z"/>`),
		g.Group(children),
	)
}