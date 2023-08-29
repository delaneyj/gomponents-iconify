package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 16.25a.75.75 0 0 1 .75.75v2c0 .138.112.25.25.25h12a.25.25 0 0 0 .25-.25v-2a.75.75 0 0 1 1.5 0v2A1.75 1.75 0 0 1 18 20.75H6A1.75 1.75 0 0 1 4.25 19v-2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path fill="currentColor" d="M10.738 3.75a.992.992 0 0 0-.988.906a36.618 36.618 0 0 0-.082 5.27c-.247.013-.493.03-.74.047l-1.49.109a.76.76 0 0 0-.585 1.167a15.555 15.555 0 0 0 4.032 4.258l.597.429a.888.888 0 0 0 1.036 0l.597-.429a15.556 15.556 0 0 0 4.032-4.258a.76.76 0 0 0-.585-1.167l-1.49-.109a42.274 42.274 0 0 0-.74-.047a36.62 36.62 0 0 0-.081-5.27a.992.992 0 0 0-.989-.906h-2.524Z"/>`),
		g.Group(children),
	)
}