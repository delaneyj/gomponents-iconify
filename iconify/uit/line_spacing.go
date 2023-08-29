package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 8h12a.5.5 0 0 0 0-1h-12a.5.5 0 0 0 0 1zm-3.459 8.17l-.874 1V6.83l.874 1a.497.497 0 0 0 .705.046a.5.5 0 0 0 .047-.705l-1.75-2a.516.516 0 0 0-.752 0l-1.75 2a.5.5 0 1 0 .752.658l.874-.998v10.338l-.874-.998a.5.5 0 0 0-.752.658l1.75 2a.499.499 0 0 0 .752 0l1.75-2a.5.5 0 0 0-.752-.658zM21.5 17h-12a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1zm0-5h-12a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}