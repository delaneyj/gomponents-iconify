package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="10.52" cy="10.52" r="1.43" fill="currentColor" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31.93 19.2L17.33 4.6a2 2 0 0 0-1.41-.6H6a2 2 0 0 0-2 2v9.92a2 2 0 0 0 .59 1.41l14.6 14.6a2 2 0 0 0 2.83 0l9.9-9.9a2 2 0 0 0 .01-2.83ZM20.62 30.52L6 15.91V6h9.92l14.6 14.62Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}