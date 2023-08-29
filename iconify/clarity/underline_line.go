package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnderlineLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 28.17c5.08 0 8.48-3.08 8.48-9V8.54a1.15 1.15 0 1 0-2.3 0v10.8c0 4.44-2.38 6.71-6.13 6.71s-6.21-2.47-6.21-6.85V8.54a1.15 1.15 0 1 0-2.3 0v10.8C9.53 25.09 13 28.17 18 28.17Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31 30H5a1.11 1.11 0 0 0 0 2.21h26A1.11 1.11 0 0 0 31 30Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}