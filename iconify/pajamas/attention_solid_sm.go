package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttentionSolidSm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.004 8L7.538 3.31A1 1 0 0 0 6.813 3H5.732a1 1 0 0 0-.866 1.5L6.886 8l-2.02 3.5a1 1 0 0 0 .866 1.5h1.081a1 1 0 0 0 .725-.31L12.004 8Z"/>`),
		g.Group(children),
	)
}