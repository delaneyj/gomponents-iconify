package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m110 319l121-121q25 25 25 60.5T231 319t-60.5 25t-60.5-25zM299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h256zM128 45q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5zm-64 0q-9 0-15 6.5t-6 15t6 15T64 88t15-6.5t6-15t-6-15T64 45zm107 342q53 0 90.5-37.5T299 259t-37.5-90.5T171 131t-90.5 37.5T43 259t37.5 90.5T171 387z"/>`),
		g.Group(children),
	)
}