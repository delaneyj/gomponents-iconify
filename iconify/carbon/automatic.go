package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automatic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 16h-4a2.002 2.002 0 0 0-2 2v12h2v-5h4v5h2V18a2.002 2.002 0 0 0-2-2zm-4 7v-5h4v5zm-6 4a10.986 10.986 0 0 1-9.216-5H12v-2H4v8h2v-3.685A13.024 13.024 0 0 0 16 29zm4-17h5.215A10.997 10.997 0 0 0 5 16H3a13.005 13.005 0 0 1 23-8.315V4h2v8h-8z"/>`),
		g.Group(children),
	)
}