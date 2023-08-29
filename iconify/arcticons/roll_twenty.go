package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.652 34.75v-21.5L24 2.5L5.348 13.25v21.5L24 45.5l18.652-10.75z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.312 32.453l6.34-19.203L24 10.585L5.348 13.25l6.34 19.203L24 45.5l12.312-13.047z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.312 32.453H11.688L24 10.585l12.312 21.868zM24 10.585V2.5m12.312 29.953l6.34 2.297m-30.964-2.297l-6.34 2.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.643 25.035c0 .827.643 1.47 1.47 1.47s1.47-.643 1.47-1.47v-1.47c0-.827-.643-1.47-1.47-1.47c-.735 0-1.47.643-1.47 1.47v1.47Zm-4.226-1.471c0-.826.643-1.47 1.47-1.47s1.47.644 1.47 1.47c0 .368-.184.735-.46 1.011c-.643.46-2.48 1.93-2.48 1.93h2.94"/>`),
		g.Group(children),
	)
}