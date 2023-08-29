package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3h298zM213 344v-43h-42v43h42zm44-165q20-20 20-48q0-36-25-61t-60-25t-60 25t-25 61h42q0-18 12.5-30.5T192 88t30.5 12.5T235 131q0 17-13 30l-26 27q-25 25-25 60v11h42q0-22 6-34.5t19-26.5z"/>`),
		g.Group(children),
	)
}