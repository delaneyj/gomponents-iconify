package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M21.282 19.992L31.234 4.73a.907.907 0 0 0-.001-1.285c-.332-.332-.949-.332-1.284.001L20 18.708L10.047 3.445c-.335-.331-.946-.333-1.282.001a.897.897 0 0 0-.265.642c0 .243.094.47.265.639l9.952 15.264l-9.95 15.265a.907.907 0 0 0 .001 1.285a.928.928 0 0 0 1.282-.003L20 21.277l9.952 15.264c.172.171.4.263.641.263a.904.904 0 0 0 .642-1.546l-9.953-15.266z"/>`),
		g.Group(children),
	)
}