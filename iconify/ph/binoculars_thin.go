package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinocularsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.22 148.09L189.6 53.41a3.94 3.94 0 0 0-.83-1.22a28 28 0 0 0-39.6 0A4 4 0 0 0 148 55v29h-40V55a4 4 0 0 0-1.17-2.83a28 28 0 0 0-39.6 0a3.94 3.94 0 0 0-.83 1.22l-41.62 94.7A44 44 0 1 0 108 168V92h40v76a44 44 0 1 0 83.22-19.91ZM64 204a36 36 0 1 1 36-36a36 36 0 0 1-36 36Zm0-80a43.78 43.78 0 0 0-22.66 6.3L73.4 57.35a20 20 0 0 1 26.6-.59v86A44 44 0 0 0 64 124Zm92-67.23a20 20 0 0 1 26.6.59l32.06 72.94A43.92 43.92 0 0 0 156 142.74ZM192 204a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}