package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m31.42 9.09l-13-6a1 1 0 0 0-.84 0l-13 6A1 1 0 0 0 4 10v17a1 1 0 0 0 .58.91l13 6a1 1 0 0 0 .84 0l13-6A1 1 0 0 0 32 27V10a1 1 0 0 0-.58-.91ZM18 5.1L28.61 10L18 14.9L7.39 10ZM6 11.56l11 5.08v14.8L6 26.36Zm13 19.88v-14.8l11-5.08v14.8Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}