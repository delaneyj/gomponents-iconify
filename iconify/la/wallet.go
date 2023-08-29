package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.969 4a2.177 2.177 0 0 0-.469.063L6.25 8.343A3.024 3.024 0 0 0 4 11.25V25c0 1.645 1.355 3 3 3h18c1.645 0 3-1.355 3-3V12c0-1.645-1.355-3-3-3H11.625L23 6v2h2V6c0-1.125-.957-2.016-2.031-2zM7 11h18c.566 0 1 .434 1 1v13c0 .566-.434 1-1 1H7c-.566 0-1-.434-1-1V12c0-.566.434-1 1-1zm15.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3z"/>`),
		g.Group(children),
	)
}