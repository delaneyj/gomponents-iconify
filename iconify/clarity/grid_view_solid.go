package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridViewSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<rect width="12" height="12" x="4" y="4" fill="currentColor" class="clr-i-solid clr-i-solid-path-1" rx="2" ry="2"/><rect width="12" height="12" x="20" y="4" fill="currentColor" class="clr-i-solid clr-i-solid-path-2" rx="2" ry="2"/><rect width="12" height="12" x="4" y="20" fill="currentColor" class="clr-i-solid clr-i-solid-path-3" rx="2" ry="2"/><rect width="12" height="12" x="20" y="20" fill="currentColor" class="clr-i-solid clr-i-solid-path-4" rx="2" ry="2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}