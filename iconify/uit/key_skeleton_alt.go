package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeletonAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.853 2.159a.5.5 0 0 0-.707-.013L15.714 7.58c-.007.006-.017.009-.024.016c-.006.006-.008.015-.014.022l-6.382 6.38A4.458 4.458 0 0 0 6.5 13a4.5 4.5 0 1 0 4.5 4.5a4.457 4.457 0 0 0-.999-2.795l6.05-6.05l1.767 1.768a.5.5 0 0 0 .707-.707l-1.767-1.768l2.122-2.121l1.767 1.768a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L19.587 5.12l2.266-2.267a.5.5 0 0 0 0-.694zM6.5 21a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7z"/>`),
		g.Group(children),
	)
}