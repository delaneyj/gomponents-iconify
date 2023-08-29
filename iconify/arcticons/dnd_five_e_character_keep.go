package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DndFiveECharacterKeep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5L5.348 13.25v21.5L24 45.5l18.652-10.75v-21.5L24 2.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5L11.688 15.547L5.348 34.75L24 37.414l18.652-2.664l-6.34-19.203L24 2.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.414L11.688 15.547h24.624L24 37.414zm0 0V45.5M11.688 15.547l-6.34-2.297m30.964 2.297l6.34-2.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.643 23.724c0 .827.643 1.47 1.47 1.47s1.47-.643 1.47-1.47v-1.47c0-.827-.643-1.47-1.47-1.47c-.735 0-1.47.643-1.47 1.47v1.47Zm-4.226-1.47c0-.827.643-1.47 1.47-1.47s1.47.643 1.47 1.47c0 .368-.184.735-.46 1.011c-.643.46-2.48 1.93-2.48 1.93h2.94"/>`),
		g.Group(children),
	)
}