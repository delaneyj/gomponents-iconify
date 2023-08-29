package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20V2Zm-2 14.5V20H4v-3.5h3.67A4.997 4.997 0 0 0 12 19a4.997 4.997 0 0 0 4.33-2.5H20Zm-16-2V4h16v10.5h-4.965l-.253.625a3.001 3.001 0 0 1-5.564 0l-.253-.625H4Zm8-9.414L7.586 9.5L9 10.914l2-2V14h2V8.914l2 2L16.414 9.5L12 5.086Z"/>`),
		g.Group(children),
	)
}