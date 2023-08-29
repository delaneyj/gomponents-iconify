package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.083 15.5a8.596 8.596 0 0 1-8.583 8.583A8.597 8.597 0 0 1 6.915 15.5A8.596 8.596 0 0 1 15.5 6.915c1.913 0 3.665.63 5.09 1.686l-1.782 1.784l8.43 2.256l-2.26-8.427l-1.89 1.89A12.036 12.036 0 0 0 15.5 3.415C8.826 3.418 3.418 8.825 3.416 15.5c.002 6.675 5.41 12.083 12.084 12.083S27.583 22.175 27.583 15.5h-3.5z"/>`),
		g.Group(children),
	)
}