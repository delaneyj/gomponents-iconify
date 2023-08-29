package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSymlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 0a1 1 0 0 0-1 1v6a.75.75 0 0 0 1.5 0V1.5H8v2.75C8 5.216 8.784 6 9.75 6h2.75v8.5H4.75a.75.75 0 0 0 0 1.5H13a1 1 0 0 0 1-1V4.476a1 1 0 0 0-.28-.695L10.362.305A1 1 0 0 0 9.643 0H3Zm9.328 4.5L9.5 1.57v2.68c0 .138.112.25.25.25h2.578ZM1 13.5a3.75 3.75 0 0 1 3.75-3.75H7V8.5c0-.445.503-.667.798-.353L10 10.5l-2.202 2.353c-.295.314-.798.091-.798-.353v-1.25H4.75A2.25 2.25 0 0 0 2.5 13.5v.75a.75.75 0 0 1-1.5 0v-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}