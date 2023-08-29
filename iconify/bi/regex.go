package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Regex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.05 3.05a7 7 0 0 0 0 9.9a.5.5 0 0 1-.707.707a8 8 0 0 1 0-11.314a.5.5 0 1 1 .707.707Zm9.9-.707a.5.5 0 0 1 .707 0a8 8 0 0 1 0 11.314a.5.5 0 0 1-.707-.707a7 7 0 0 0 0-9.9a.5.5 0 0 1 0-.707ZM6 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm5-6.5a.5.5 0 0 0-1 0v2.117L8.257 5.57a.5.5 0 0 0-.514.858L9.528 7.5L7.743 8.571a.5.5 0 1 0 .514.858L10 8.383V10.5a.5.5 0 1 0 1 0V8.383l1.743 1.046a.5.5 0 0 0 .514-.858L11.472 7.5l1.785-1.071a.5.5 0 1 0-.514-.858L11 6.617V4.5Z"/>`),
		g.Group(children),
	)
}