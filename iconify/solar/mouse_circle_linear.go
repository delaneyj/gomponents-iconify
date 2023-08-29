package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M8 10a4 4 0 1 1 8 0v4a4 4 0 0 1-8 0v-4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8.5 10h7M12 10V6"/><path fill="currentColor" d="m10.013 2.033l-.15-.735l.15.735Zm6.681.135a.75.75 0 0 0-.66 1.347l.66-1.347Zm-6.83-.87c-4.918 1-8.614 5.377-8.614 10.619h1.5c0-4.524 3.188-8.29 7.412-9.15l-.299-1.47ZM1.25 11.917C1.25 17.894 6.057 22.75 12 22.75v-1.5c-5.103 0-9.25-4.173-9.25-9.333h-1.5ZM12 22.75c5.943 0 10.75-4.856 10.75-10.833h-1.5c0 5.16-4.147 9.333-9.25 9.333v1.5Zm-.75-18.9v5.042h1.5V3.85h-1.5Zm11.5 8.067c0-4.286-2.47-7.992-6.056-9.749l-.66 1.347a9.341 9.341 0 0 1 5.216 8.402h1.5Zm-12.588-9.15c.508-.103 1.088.324 1.088 1.083h1.5c0-1.468-1.229-2.89-2.887-2.552l.3 1.47Z"/></g>`),
		g.Group(children),
	)
}