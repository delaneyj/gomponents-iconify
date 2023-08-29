package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3h298zM192 73q-24 0-41 17t-17 41t17 40.5t41 16.5t41-16.5t17-40.5t-17-41t-41-17zm128 228v-19q0-20-23.5-35.5t-52.5-23t-52-7.5t-52 7.5t-52.5 23T64 282v19h256z"/>`),
		g.Group(children),
	)
}