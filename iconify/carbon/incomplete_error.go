package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncompleteError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 26a12 12 0 0 1 0-24zm3.826-21.236a10.029 10.029 0 0 1 3.242 2.168L22.48 5.52a12.036 12.036 0 0 0-3.89-2.602zM26 14a11.93 11.93 0 0 0-.917-4.59l-1.847.764A9.943 9.943 0 0 1 24 14zm4 10a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 24 20a4.005 4.005 0 0 1 4 4zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 24 28a4.005 4.005 0 0 1-4-4z"/>`),
		g.Group(children),
	)
}