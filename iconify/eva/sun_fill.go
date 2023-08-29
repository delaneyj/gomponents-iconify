package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSunFill0"><g id="evaSunFill1"><path id="evaSunFill2" fill="currentColor" d="M12 6a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM6 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1Zm.22-7a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0 0-1.41ZM17 8.14a1 1 0 0 0 .69-.28l1.44-1.39A1 1 0 0 0 17.78 5l-1.44 1.42a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.31ZM12 18a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm5.73-1.86a1 1 0 0 0-1.39 1.44L17.78 19a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.3a1 1 0 0 0 0-1.42Zm-11.46 0l-1.44 1.39a1 1 0 0 0 0 1.42a1 1 0 0 0 .72.3a1 1 0 0 0 .67-.25l1.44-1.39a1 1 0 0 0-1.39-1.44ZM12 8a4 4 0 1 0 4 4a4 4 0 0 0-4-4Z"/></g></g>`),
		g.Group(children),
	)
}