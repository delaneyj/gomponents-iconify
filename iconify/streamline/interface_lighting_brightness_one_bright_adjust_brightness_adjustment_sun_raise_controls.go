package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLightingBrightnessOneBrightAdjustBrightnessAdjustmentSunRaiseControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="M7 .5v2m-4.6-.1l1.42 1.42M.5 7h2m-.1 4.6l1.42-1.42M7 13.5v-2m4.6.1l-1.42-1.42M13.5 7h-2m.1-4.6l-1.42 1.42"/></g>`),
		g.Group(children),
	)
}