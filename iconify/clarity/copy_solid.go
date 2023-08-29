package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M27 3.56A1.56 1.56 0 0 0 25.43 2H5.57A1.56 1.56 0 0 0 4 3.56v24.88A1.56 1.56 0 0 0 5.57 30h.52V4.07H27Z" class="clr-i-solid clr-i-solid-path-1"/><rect width="23" height="28" x="8" y="6" fill="currentColor" class="clr-i-solid clr-i-solid-path-2" rx="1.5" ry="1.5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}