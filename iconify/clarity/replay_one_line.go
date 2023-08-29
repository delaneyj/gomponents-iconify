package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplayOneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M19 27.27a1 1 0 0 0 1-1V14a1 1 0 0 0-1-1a3.8 3.8 0 0 0-1.1.23l-2 .62a.92.92 0 0 0-.72.86a.88.88 0 0 0 .88.86a1.46 1.46 0 0 0 .43-.08l1.51-.42v11.2a1 1 0 0 0 1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18.06 5h-6.7l2.92-2.64A1 1 0 0 0 12.94.88L7.32 6l5.62 5a1 1 0 0 0 .67.26a1 1 0 0 0 .74-.33a1 1 0 0 0-.07-1.42L11.46 7h6.6A11.78 11.78 0 1 1 7.71 24.41a1 1 0 0 0-1.71.95A13.78 13.78 0 1 0 18.06 5Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}