package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M10.65 24.44a9.51 9.51 0 0 1 7.06-9.17L13 3a1 1 0 0 0-1.87 0L2.07 26.56A1 1 0 0 0 3 27.92h8.32a9.44 9.44 0 0 1-.67-3.48Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M32 10H20a1 1 0 0 0-1 1v4a9.43 9.43 0 0 1 10.63 9H32a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1Z" class="clr-i-solid clr-i-solid-path-2"/><circle cx="20.15" cy="24.44" r="7.5" fill="currentColor" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}