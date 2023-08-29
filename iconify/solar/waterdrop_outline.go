package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterdropOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.7 6.957a.75.75 0 1 0-.629-1.361c-1.965.908-3.48 2.715-4.171 4.901a.75.75 0 0 0 1.43.453c.577-1.825 1.825-3.278 3.37-3.993Z"/><path fill-rule="evenodd" d="M14.386 1.729a6.184 6.184 0 0 0-4.772 0C5.154 3.59 2.25 8.022 2.25 12.929v.264c0 5.262 4.175 9.557 9.361 9.557h.777c5.187 0 9.362-4.295 9.362-9.557v-.265c0-4.906-2.903-9.337-7.364-11.2Zm-4.194 1.384a4.684 4.684 0 0 1 3.616 0c3.888 1.623 6.442 5.498 6.442 9.815v.265c0 4.466-3.535 8.057-7.861 8.057h-.777c-4.327 0-7.862-3.591-7.862-8.057v-.265c0-4.317 2.554-8.192 6.442-9.815Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}