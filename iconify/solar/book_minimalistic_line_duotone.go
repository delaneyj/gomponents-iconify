package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M4 8c0-2.828 0-4.243.879-5.121C5.757 2 7.172 2 10 2h4c2.828 0 4.243 0 5.121.879C20 3.757 20 5.172 20 8v8c0 2.828 0 4.243-.879 5.121C18.243 22 16.828 22 14 22h-4c-2.828 0-4.243 0-5.121-.879C4 20.243 4 18.828 4 16V8Z"/><path fill="currentColor" d="m6.121 16.102l-.194-.724l.194.724ZM3.276 18.03a.75.75 0 0 0 1.448.389l-1.448-.389ZM6.25 16a.75.75 0 0 0 1.5 0h-1.5Zm1.5-13.5a.75.75 0 0 0-1.5 0h1.5Zm.148 14.25h12v-1.5h-12v1.5Zm0-1.5c-.882 0-1.47-.006-1.97.128l.387 1.449c.263-.07.605-.077 1.583-.077v-1.5Zm-1.97.128a3.75 3.75 0 0 0-2.652 2.651l1.448.389a2.25 2.25 0 0 1 1.591-1.591l-.388-1.45ZM7.75 16V2.5h-1.5V16h1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}