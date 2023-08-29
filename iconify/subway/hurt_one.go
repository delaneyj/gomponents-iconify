package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HurtOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 29.8c-64 0-96.2 27.6-128 64c-31.8-36.4-64-64-128-64S0 72.5 0 200.5c0 64 64 192 256 298.7c192-106.7 256-234.7 256-298.7c0-128-64-170.7-128-170.7zM256 450C81.7 346.6 42.7 235.7 42.7 200.5c0-58.4 14.8-128 85.3-128c44.8 0 66.6 15.9 95.9 49.4l32.1 35.9l32.1-35.9c29.3-33.5 51.1-49.4 95.9-49.4c70.5 0 85.3 69.6 85.3 128c0 35.2-39 146.1-213.3 249.5z"/>`),
		g.Group(children),
	)
}