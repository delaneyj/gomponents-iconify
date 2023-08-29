package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m158 184l-1 2L79 51Q137 3 213 3q23 0 47 5zm259-32H211l78-135q46 17 79.5 52.5T417 152zm5 21q5 22 5 43q0 83-57 144L269 184l-6-11h159zm-282 43l24 43H4q-4-22-4-43q0-82 56-144zM10 280h206l-78 135q-46-17-79.5-52.5T10 280zm240 0l20-34l78 135q-59 48-135 48q-22 0-46-5z"/>`),
		g.Group(children),
	)
}