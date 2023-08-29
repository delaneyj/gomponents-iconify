package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.25 3.85c0-.759-.58-1.185-1.088-1.082c-4.224.859-7.412 4.626-7.412 9.15c0 5.16 4.147 9.332 9.25 9.332s9.25-4.172 9.25-9.333c0-3.696-2.13-6.89-5.216-8.401a.75.75 0 1 1 .66-1.347a10.841 10.841 0 0 1 6.056 9.748c0 5.977-4.807 10.833-10.75 10.833S1.25 17.894 1.25 11.917c0-5.242 3.696-9.62 8.613-10.619c1.658-.337 2.887 1.084 2.887 2.553v1.458c2.267.36 4 2.323 4 4.691v4a4.75 4.75 0 1 1-9.5 0v-4a4.751 4.751 0 0 1 4-4.69V3.85Zm0 2.987A3.255 3.255 0 0 0 8.837 9.25h2.413V6.837Zm1.5 2.413h2.413a3.255 3.255 0 0 0-2.413-2.413V9.25Zm2.5 1.5h-6.5V14a3.25 3.25 0 1 0 6.5 0v-3.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}