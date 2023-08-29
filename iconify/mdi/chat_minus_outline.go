package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3c5.5 0 10 3.58 10 8c0 .58-.08 1.14-.22 1.68c-.59-.3-1.23-.52-1.9-.62c.08-.34.12-.7.12-1.06c0-3.31-3.58-6-8-6s-8 2.69-8 6s3.58 6 8 6l1.09-.05L13 18l.08.95L12 19c-1.19 0-2.38-.17-3.53-.5C6.64 20 4.37 20.89 2 21c2.33-2.33 2.75-3.9 2.75-4.5A7.218 7.218 0 0 1 2 11c0-4.42 4.5-8 10-8m11 14v2h-8v-2h8Z"/>`),
		g.Group(children),
	)
}