package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPlayCircleFill0"><g id="evaPlayCircleFill1"><g id="evaPlayCircleFill2" fill="currentColor"><path d="m11.5 14.6l2.81-2.6l-2.81-2.6v5.2z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm4 11.18l-3.64 3.37a1.74 1.74 0 0 1-1.16.45a1.68 1.68 0 0 1-.69-.15a1.6 1.6 0 0 1-1-1.48V8.63a1.6 1.6 0 0 1 1-1.48a1.7 1.7 0 0 1 1.85.3L16 10.82a1.6 1.6 0 0 1 0 2.36Z"/></g></g></g>`),
		g.Group(children),
	)
}