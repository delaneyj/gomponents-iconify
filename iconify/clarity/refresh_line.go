package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M22.4 11.65a1.09 1.09 0 0 0 1.09 1.09h10.94V1.81a1.09 1.09 0 1 0-2.19 0v7.14a16.41 16.41 0 1 0 1.47 15.86a1.12 1.12 0 0 0-2.05-.9a14.18 14.18 0 1 1-1.05-13.36H23.5a1.09 1.09 0 0 0-1.1 1.1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}