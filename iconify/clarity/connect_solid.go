package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M17 12h-2.85a6.25 6.25 0 0 0-6.21 5H2v2h5.93a6.22 6.22 0 0 0 6.22 5H17Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M28.23 17A6.25 6.25 0 0 0 22 12h-3v12h3a6.22 6.22 0 0 0 6.22-5H34v-2Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}