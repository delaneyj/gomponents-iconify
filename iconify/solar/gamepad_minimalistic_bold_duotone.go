package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13.172 8c2.539 0 3.808 0 4.784.614c.247.155.477.337.684.542c.82.81 1.109 2.047 1.687 4.52l1.022 4.373A2.858 2.858 0 0 1 16 19.957l-.12-.246a3.203 3.203 0 0 0-2.877-1.794h-2.004A3.203 3.203 0 0 0 8.12 19.71l-.12.246a2.858 2.858 0 0 1-5.35-1.908l1.022-4.374c.578-2.472.867-3.708 1.687-4.519a4 4 0 0 1 .684-.542C7.02 8 8.289 8 10.828 8h2.344Z" opacity=".5"/><path d="M9.25 12a.75.75 0 0 0-1.5 0v.75H7a.75.75 0 0 0 0 1.5h.75V15a.75.75 0 0 0 1.5 0v-.75H10a.75.75 0 0 0 0-1.5h-.75V12ZM15 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM15.75 4a.75.75 0 0 0-1.5 0v1a.25.25 0 0 1-.25.25h-1A1.75 1.75 0 0 0 11.25 7v1h1.5V7a.25.25 0 0 1 .25-.25h1A1.75 1.75 0 0 0 15.75 5V4Z"/></g>`),
		g.Group(children),
	)
}