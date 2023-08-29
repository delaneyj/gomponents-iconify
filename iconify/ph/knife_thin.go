package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M229 35a24 24 0 0 0-34 0L21.14 213.4a4 4 0 0 0 2 6.7a150.65 150.65 0 0 0 34 3.9c32.64 0 65.34-10.65 96.12-31.52c31.2-21.15 49.21-45 50-46a4 4 0 0 0-.37-5.24l-21.48-21.5L229 69a24.09 24.09 0 0 0 0-34Zm-34.35 109.41a226.48 226.48 0 0 1-45.89 41.44c-37.93 25.72-77.18 35.09-116.76 27.9L146.53 96.24Zm28.68-81.07a.76.76 0 0 0-.09.09l-47.54 50.68l-23.58-23.6l48.57-49.83a16 16 0 0 1 22.64 22.66Z"/>`),
		g.Group(children),
	)
}