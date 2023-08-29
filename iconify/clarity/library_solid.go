package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibrarySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M12.75 3h-7.5A1.15 1.15 0 0 0 4 4v29h10V4a1.15 1.15 0 0 0-1.25-1Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="m33.77 31.09l-6.94-18.3a1 1 0 0 0-1.29-.58L22 13.59V9a1 1 0 0 0-1-1h-5v25h6V14.69L28.93 33Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}