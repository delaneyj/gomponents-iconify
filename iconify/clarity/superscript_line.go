package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m14.43 18l6.79 8.6a1.17 1.17 0 0 1-.92 1.9a1.17 1.17 0 0 1-.92-.44l-6.44-8.13L6.47 28a1.17 1.17 0 0 1-.92.44a1.17 1.17 0 0 1-.92-1.9l6.8-8.54l-6.8-8.6a1.17 1.17 0 0 1 .92-1.9a1.2 1.2 0 0 1 .96.5l6.43 8.13L19.38 8a1.17 1.17 0 0 1 .92-.44a1.17 1.17 0 0 1 .92 1.9Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m22.85 14.47l4.51-3.85a9.37 9.37 0 0 0 1.88-2a3.43 3.43 0 0 0 .59-1.86a2.27 2.27 0 0 0-.36-1.27a2.38 2.38 0 0 0-.95-.83a2.77 2.77 0 0 0-1.26-.29a3.39 3.39 0 0 0-1.83.5a5.83 5.83 0 0 0-1.49 1.42l-1-.81a5.12 5.12 0 0 1 4.36-2.37a4.36 4.36 0 0 1 2 .45a3.47 3.47 0 0 1 2 3.18A4.44 4.44 0 0 1 30.58 9a11.14 11.14 0 0 1-2.24 2.46l-3.24 2.85h6.28v1.33h-8.53Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}