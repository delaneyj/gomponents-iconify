package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M178 78A84 84 0 1 0 78 178A84 84 0 1 0 178 78Zm42 82a60.75 60.75 0 0 1-.38 6.65l-44-44a83.62 83.62 0 0 0 4-19.39A59.83 59.83 0 0 1 220 160ZM36 96a60 60 0 1 1 60 60a60.07 60.07 0 0 1-60-60Zm67.28 83.66a83.62 83.62 0 0 0 19.39-4l44 44a60.75 60.75 0 0 1-6.67.34a59.83 59.83 0 0 1-56.72-40.34Zm88.53 31.18l-46.73-46.73a85 85 0 0 0 19-19l46.73 46.73a60.45 60.45 0 0 1-19 19Z"/>`),
		g.Group(children),
	)
}