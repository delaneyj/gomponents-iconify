package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResourcePoolOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.2 13.4a13.91 13.91 0 0 1 .75 3.6H17.49L8.3 28.07A14 14 0 0 1 22.61 4.8a7.43 7.43 0 0 1 .58-1.92a16.06 16.06 0 1 0 9.93 9.93a7.43 7.43 0 0 1-1.92.59ZM18 32a13.91 13.91 0 0 1-8.16-2.65L18.43 19h13.52A14 14 0 0 1 18 32Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}