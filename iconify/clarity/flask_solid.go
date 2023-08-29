package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.49 27.4L23 14.94V4h1a1 1 0 0 0 0-2H12.08a1 1 0 0 0 0 2H13v10.94L4.58 27.31a4.31 4.31 0 0 0-.78 3A4.23 4.23 0 0 0 8 34h19.86A4.36 4.36 0 0 0 31 32.8a4.23 4.23 0 0 0 .49-5.4ZM15 15.49V4h6v11.49L26.15 23H9.85Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}