package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33.16 19.13L19.58 5.55A1.92 1.92 0 0 0 18.21 5h-2.09l15.63 15.45l-10.53 10.62a1.93 1.93 0 0 0 2.73 0l9.21-9.21a1.93 1.93 0 0 0 0-2.73Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M27.78 19.17L14.2 5.58A1.92 1.92 0 0 0 12.83 5H3.61a1.93 1.93 0 0 0-1.93 1.93v9.22a1.92 1.92 0 0 0 .57 1.36L15.84 31.1a1.93 1.93 0 0 0 2.73 0l9.21-9.21a1.93 1.93 0 0 0 0-2.72ZM6.67 11.72A1.73 1.73 0 1 1 8.4 10a1.73 1.73 0 0 1-1.73 1.72Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}