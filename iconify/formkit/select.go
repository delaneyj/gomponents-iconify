package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Select(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M40 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h36a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1ZM4 0a4 4 0 0 0-4 4v16a4 4 0 0 0 4 4h36a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4Zm4 11.5c0-.828.38-1.5.848-1.5h11.304c.468 0 .848.672.848 1.5s-.38 1.5-.848 1.5H8.848C8.38 13 8 12.328 8 11.5Zm23.707-1.207L31 9.586L29.586 11l.707.707l2 2l.707.707l.707-.707l2-2l.707-.707L35 9.586l-.707.707L33 11.586l-1.293-1.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}