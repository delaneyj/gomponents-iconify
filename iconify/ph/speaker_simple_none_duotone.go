package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSimpleNoneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M152 32v192l-72-56H32a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h48Z" opacity=".2"/><path d="M155.51 24.81a8 8 0 0 0-8.42.88L77.25 80H32a16 16 0 0 0-16 16v64a16 16 0 0 0 16 16h45.25l69.84 54.31A8 8 0 0 0 160 224V32a8 8 0 0 0-4.49-7.19ZM144 207.64l-59.09-45.95A7.94 7.94 0 0 0 80 160H32V96h48a7.94 7.94 0 0 0 4.91-1.69L144 48.36Z"/></g>`),
		g.Group(children),
	)
}