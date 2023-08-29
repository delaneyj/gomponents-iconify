package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMoonFill0"><g id="evaMoonFill1"><path id="evaMoonFill2" fill="currentColor" d="M12.3 22h-.1a10.31 10.31 0 0 1-7.34-3.15a10.46 10.46 0 0 1-.26-14a10.13 10.13 0 0 1 4-2.74a1 1 0 0 1 1.06.22a1 1 0 0 1 .24 1a8.4 8.4 0 0 0 1.94 8.81a8.47 8.47 0 0 0 8.83 1.94a1 1 0 0 1 1.27 1.29A10.16 10.16 0 0 1 19.6 19a10.28 10.28 0 0 1-7.3 3Z"/></g></g>`),
		g.Group(children),
	)
}