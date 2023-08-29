package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 16V9C42 7.34315 40.6569 6 39 6H9C7.34315 6 6 7.34315 6 9V16"/><path d="M6 32V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V32"/><path d="M5 24H13.075L20 16L27 32L33.975 24H43"/></g>`),
		g.Group(children),
	)
}