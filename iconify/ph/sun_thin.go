package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124 40V16a4 4 0 0 1 8 0v24a4 4 0 0 1-8 0Zm64 88a60 60 0 1 1-60-60a60.07 60.07 0 0 1 60 60Zm-8 0a52 52 0 1 0-52 52a52.06 52.06 0 0 0 52-52ZM61.17 66.83a4 4 0 0 0 5.66-5.66l-16-16a4 4 0 0 0-5.66 5.66Zm0 122.34l-16 16a4 4 0 0 0 5.66 5.66l16-16a4 4 0 0 0-5.66-5.66ZM192 68a4 4 0 0 0 2.83-1.17l16-16a4 4 0 1 0-5.66-5.66l-16 16A4 4 0 0 0 192 68Zm2.83 121.17a4 4 0 0 0-5.66 5.66l16 16a4 4 0 0 0 5.66-5.66ZM40 124H16a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Zm88 88a4 4 0 0 0-4 4v24a4 4 0 0 0 8 0v-24a4 4 0 0 0-4-4Zm112-88h-24a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}