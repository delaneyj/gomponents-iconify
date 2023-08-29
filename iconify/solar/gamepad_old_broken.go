package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadOldBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#solarGamepadOldBroken0)"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8 10v4m-2-2h4"/><path fill="currentColor" d="M16 10.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 4V3a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1V0m7 12c0 2.8 0 4.2-.545 5.27a5 5 0 0 1-2.185 2.185C18.2 20 16.8 20 14 20h-4c-2.8 0-4.2 0-5.27-.545a5 5 0 0 1-2.185-2.185C2 16.2 2 14.8 2 12c0-2.8 0-4.2.545-5.27A5 5 0 0 1 4.73 4.545C5.8 4 7.2 4 10 4h4c2.8 0 4.2 0 5.27.545a5 5 0 0 1 2.185 2.185c.188.37.311.778.392 1.27"/></g><defs><clipPath id="solarGamepadOldBroken0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}