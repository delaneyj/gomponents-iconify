package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSyncTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 5.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 0 0-9 0Zm6.5-3a.5.5 0 0 1 .5.5v1.5a.5.5 0 0 1-.5.5H15a.5.5 0 0 1 0-1h.468a1.982 1.982 0 0 0-.933-.25a2 2 0 0 0-1.45.586a.5.5 0 1 1-.706-.707A3 3 0 0 1 16 3.152V3a.5.5 0 0 1 .5-.5Zm-.876 5.532A3 3 0 0 1 13 7.848V8a.5.5 0 0 1-1 0V6.5a.5.5 0 0 1 .5-.5H14a.5.5 0 0 1 0 1h-.468a1.982 1.982 0 0 0 .933.25a2 2 0 0 0 1.45-.586a.5.5 0 1 1 .706.707a3 3 0 0 1-.997.66ZM2.5 7h6.707c.099.349.23.683.393 1H2.5a.5.5 0 0 1-.09-.992L2.5 7Zm9 8a.5.5 0 0 1 .09.992L11.5 16h-5a.5.5 0 0 1-.09-.992L6.5 15h5Zm-7-4h9a.5.5 0 0 1 .09.992L13.5 12h-9a.5.5 0 0 1-.09-.992L4.5 11Z"/>`),
		g.Group(children),
	)
}