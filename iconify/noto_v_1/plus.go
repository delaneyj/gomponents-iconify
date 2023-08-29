package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#ed6c30" d="M127.35 59.61c0-4.36-3.56-7.92-7.92-7.92h-35.2c-4.36 0-7.92-3.56-7.92-7.92V8.57c0-4.36-3.56-7.92-7.92-7.92h-8.78c-4.36 0-7.92 3.56-7.92 7.92v35.19c0 4.36-3.56 7.92-7.92 7.92H8.57c-4.36 0-7.92 3.56-7.92 7.92v8.79c0 4.36 3.56 7.92 7.92 7.92h35.2c4.36 0 7.92 3.56 7.92 7.92v35.19c0 4.36 3.56 7.92 7.92 7.92h8.78c4.36 0 7.92-3.56 7.92-7.92V84.23c0-4.36 3.56-7.92 7.92-7.92h35.2c4.36 0 7.92-3.56 7.92-7.92v-8.78z"/>`),
		g.Group(children),
	)
}