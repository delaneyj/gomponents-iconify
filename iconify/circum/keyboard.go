package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.437 18.5H4.562a2.5 2.5 0 0 1-2.5-2.5V8a2.5 2.5 0 0 1 2.5-2.5h14.875a2.5 2.5 0 0 1 2.5 2.5v8a2.5 2.5 0 0 1-2.5 2.5ZM4.562 6.5a1.5 1.5 0 0 0-1.5 1.5v8a1.5 1.5 0 0 0 1.5 1.5h14.875a1.5 1.5 0 0 0 1.5-1.5V8a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M5.548 16.5h12.9a.5.5 0 0 0 0-1h-12.9a.5.5 0 0 0 0 1Z"/><circle cx="5.82" cy="9.248" r=".75" fill="currentColor"/><circle cx="9.94" cy="9.248" r=".75" fill="currentColor"/><circle cx="14.06" cy="9.248" r=".75" fill="currentColor"/><circle cx="18.18" cy="9.248" r=".75" fill="currentColor"/><circle cx="5.82" cy="12.998" r=".75" fill="currentColor"/><circle cx="9.94" cy="12.998" r=".75" fill="currentColor"/><circle cx="14.06" cy="12.998" r=".75" fill="currentColor"/><circle cx="18.18" cy="12.998" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}