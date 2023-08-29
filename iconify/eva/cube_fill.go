package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCubeFill0"><g id="evaCubeFill1"><path id="evaCubeFill2" fill="currentColor" d="M11.25 11.83L3 8.36v7.73a1.69 1.69 0 0 0 1 1.52L11.19 21h.06ZM12 10.5l8.51-3.57a1.62 1.62 0 0 0-.51-.38l-7.2-3.37a1.87 1.87 0 0 0-1.6 0L4 6.55a1.62 1.62 0 0 0-.51.38Zm.75 1.33V21h.05l7.2-3.39a1.69 1.69 0 0 0 1-1.51V8.36Z"/></g></g>`),
		g.Group(children),
	)
}