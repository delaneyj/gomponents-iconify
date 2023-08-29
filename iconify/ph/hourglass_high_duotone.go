package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassHighDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 64v11.64a8 8 0 0 1-3.18 6.36L128 128L67.2 82.4A8 8 0 0 1 64 76V64Z" opacity=".2"/><path d="M184 24H72a16 16 0 0 0-16 16v36a16.07 16.07 0 0 0 6.4 12.8l52.27 39.2l-52.27 39.2A16.07 16.07 0 0 0 56 180v36a16 16 0 0 0 16 16h112a16 16 0 0 0 16-16v-35.64a16.09 16.09 0 0 0-6.35-12.77L141.27 128l52.38-39.59A16.09 16.09 0 0 0 200 75.64V40a16 16 0 0 0-16-16Zm0 16v16H72V40Zm0 176H72v-36l56-42l56 42.35Zm-56-98L72 76v-4h112v3.64Z"/></g>`),
		g.Group(children),
	)
}