package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M5 7h2v2H5z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M9 7h2v2H9z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13 7h2v2h-2z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M32 4H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2ZM4 6h28v4.2H4Zm0 24V11.8h28V30Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}