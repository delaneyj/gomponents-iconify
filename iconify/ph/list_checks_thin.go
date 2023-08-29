package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListChecksThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4h-88a4 4 0 0 1 0-8h88a4 4 0 0 1 4 4Zm-92-60h88a4 4 0 0 0 0-8h-88a4 4 0 0 0 0 8Zm88 120h-88a4 4 0 0 0 0 8h88a4 4 0 0 0 0-8ZM85.17 45.17L56 74.34L42.83 61.17a4 4 0 0 0-5.66 5.66l16 16a4 4 0 0 0 5.66 0l32-32a4 4 0 0 0-5.66-5.66Zm0 64L56 138.34l-13.17-13.17a4 4 0 1 0-5.66 5.66l16 16a4 4 0 0 0 5.66 0l32-32a4 4 0 0 0-5.66-5.66Zm0 64L56 202.34l-13.17-13.17a4 4 0 0 0-5.66 5.66l16 16a4 4 0 0 0 5.66 0l32-32a4 4 0 0 0-5.66-5.66Z"/>`),
		g.Group(children),
	)
}