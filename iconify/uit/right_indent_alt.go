package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightIndentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.82 9.116a.5.5 0 0 0-.64.768L4.719 12l-2.54 2.116a.5.5 0 0 0 .641.768l3-2.5a.5.5 0 0 0 0-.768l-3-2.5zM12.5 7h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1zM9.045 5h-.003a.5.5 0 0 0-.5.497l-.084 13a.5.5 0 0 0 .497.503h.003a.5.5 0 0 0 .5-.497l.084-13A.5.5 0 0 0 9.045 5zM21.5 10h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1zm0 8h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1zm0-4h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}