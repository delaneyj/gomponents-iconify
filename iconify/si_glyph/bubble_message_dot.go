package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.1 0H3.902C2.873 0 2.039.795 2.039 2.094v5.844c0 1.299.834 2.051 1.863 2.051h5.7l-1.57 4.98l4.474-4.98h1.595c1.029 0 1.864-.752 1.864-2.051V2.094C15.964.795 15.129 0 14.1 0zM5.02 6.312a1.31 1.31 0 0 1-1.301-1.318c0-.728.582-1.317 1.301-1.317c.721 0 1.301.59 1.301 1.317A1.309 1.309 0 0 1 5.02 6.312zM9 6.24c-.709 0-1.281-.558-1.281-1.246S8.291 3.749 9 3.749c.707 0 1.281.558 1.281 1.245c0 .689-.574 1.246-1.281 1.246zm4 0c-.705 0-1.281-.558-1.281-1.246S12.295 3.749 13 3.749c.707 0 1.281.558 1.281 1.245c0 .689-.574 1.246-1.281 1.246z"/>`),
		g.Group(children),
	)
}