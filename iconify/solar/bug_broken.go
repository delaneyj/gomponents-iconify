package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M14 21.71A7 7 0 0 1 5 15v-3.063A3.938 3.938 0 0 1 8.938 8h6.124A3.938 3.938 0 0 1 19 11.938V15a6.977 6.977 0 0 1-2 4.899"/><path fill="currentColor" d="M8.25 7.5a.75.75 0 0 0-1.5 0h1.5Zm-.615-2.917a.75.75 0 1 0 1.246.834l-1.246-.834ZM17.25 8.5v-1h-1.5v1h1.5Zm-10.5-1v1h1.5v-1h-1.5Zm10.5 0c0-2.9-2.35-5.25-5.25-5.25v1.5a3.75 3.75 0 0 1 3.75 3.75h1.5ZM12 2.25a5.246 5.246 0 0 0-4.365 2.333l1.246.834A3.746 3.746 0 0 1 12 3.75v-1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19 14h3M5 14H2M14.5 3.5L17 2M9.5 3.5L7 2m13.5 18l-2-.8m2-11.2l-2 .8M3.5 20l2-.8M3.5 8l2 .8M12 21.5V15"/></g>`),
		g.Group(children),
	)
}