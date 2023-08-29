package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freedombox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24l9.08-5.09l-5.73-3.13c5.3-4.86 10.49-6.93 12.65-6.93s3.49.66 3.49 2.66c0 2.86-2 3.19-2 9c0 4.64-3.1 5.84-5.21 5.72a7.66 7.66 0 1 1-12.25 9a7.66 7.66 0 1 1-12.25-9c-2.11.12-5.21-1.08-5.21-5.72c0-5.86-2-6.19-2-9c0-2 1.41-2.66 3.49-2.66s7.36 2.07 12.66 6.93l-5.73 3.13Zm0 0v8.67m-2.27-9.95v8.68m-2.27-9.95v8.67m-2.27-9.94v8.67m-2.27-9.94v8.67"/>`),
		g.Group(children),
	)
}