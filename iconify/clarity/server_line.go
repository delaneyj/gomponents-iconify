package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 24.3a2.48 2.48 0 1 0 2.48 2.47A2.48 2.48 0 0 0 18 24.3Zm0 3.6a1.13 1.13 0 1 1 1.13-1.12A1.13 1.13 0 0 1 18 27.9Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M25.65 3.6h-15.3A1.35 1.35 0 0 0 9 4.95V32.4h18V4.95a1.35 1.35 0 0 0-1.35-1.35Zm-.45 27H10.8V5.4h14.4Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M12.6 7.2h10.8v1.44H12.6z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M12.6 10.8h10.8v1.44H12.6z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}