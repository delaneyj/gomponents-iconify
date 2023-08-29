package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33.12 12.81a7.49 7.49 0 0 1-9.93-9.93a16.06 16.06 0 1 0 9.93 9.93Zm-15.34 15a1.65 1.65 0 1 1 1.65-1.65a1.65 1.65 0 0 1-1.65 1.69Zm1.37-8.06v1.72a1.37 1.37 0 0 1-1.3 1.36h-.11a1.34 1.34 0 0 1-1.39-1.3v-2.76a1.19 1.19 0 0 1 1.12-1.31c1.57-.12 4.18-.7 4.18-3.25c0-1.83-1.41-3.07-3.43-3.07a5.31 5.31 0 0 0-4 1.92a1.36 1.36 0 0 1-2.35-.9a1.43 1.43 0 0 1 .43-1a7.77 7.77 0 0 1 6-2.69c3.7 0 6.28 2.3 6.28 5.6c0 3.09-1.97 5.13-5.43 5.72Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}