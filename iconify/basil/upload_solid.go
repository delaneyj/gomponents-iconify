package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 16.25a.75.75 0 0 1 .75.75v2c0 .138.112.25.25.25h12a.25.25 0 0 0 .25-.25v-2a.75.75 0 0 1 1.5 0v2A1.75 1.75 0 0 1 18 20.75H6A1.75 1.75 0 0 1 4.25 19v-2a.75.75 0 0 1 .75-.75Zm5.738-.123a.992.992 0 0 1-.989-.905a36.618 36.618 0 0 1-.08-5.27a42.17 42.17 0 0 1-.741-.048l-1.49-.109a.76.76 0 0 1-.585-1.167a15.555 15.555 0 0 1 4.032-4.258l.597-.429a.888.888 0 0 1 1.036 0l.597.43a15.556 15.556 0 0 1 4.032 4.257a.76.76 0 0 1-.585 1.167l-1.49.109c-.246.018-.493.034-.74.047c.1 1.757.072 3.518-.082 5.27a.992.992 0 0 1-.988.906h-2.524Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}