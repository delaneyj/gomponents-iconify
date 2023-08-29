package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 4H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Zm0 6.2H4V6h28Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M5 7h2v2H5z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M9 7h2v2H9z" class="clr-i-solid clr-i-solid-path-3"/><path fill="currentColor" d="M13 7h2v2h-2z" class="clr-i-solid clr-i-solid-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}