package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.998 3v2c0 9.627-5.373 14-12 14H5.241a17.14 17.14 0 0 0-.243 3h-2c0-1.363.116-2.6.346-3.732c-.23-1.294-.346-3.05-.346-5.268c0-5.523 4.477-10 10-10c2 0 4 1 8 0Zm-8 2a8 8 0 0 0-8 8c0 .362.003.711.01 1.046c1.254-1.978 3.091-3.541 5.494-4.914l.992 1.736C8.64 12.5 6.746 14.354 5.774 17h3.224c6.015 0 9.871-3.973 9.997-11.612c-1.372.133-2.647.048-4.22-.188c-1.15-.173-1.376-.2-1.777-.2Z"/>`),
		g.Group(children),
	)
}