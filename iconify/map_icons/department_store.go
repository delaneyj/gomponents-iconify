package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DepartmentStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M34 11V7.964C34 5.483 31.891 3 29.411 3h-8.343C18.589 3 16 5.483 16 7.964V11h-5.312L5 46h40.187l-5.778-35H34zM19 7.964C19 7.035 20.14 6 21.068 6h8.343C30.338 6 31 7.035 31 7.964V11H19V7.964z"/>`),
		g.Group(children),
	)
}