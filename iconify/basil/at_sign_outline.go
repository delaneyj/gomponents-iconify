package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtSignOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.755 12a7.25 7.25 0 1 1 14.5 0c0 .034.002.067.006.1a.757.757 0 0 0-.004.095c.007.255-.026.65-.11 1.027c-.09.411-.21.647-.279.72a1.65 1.65 0 0 1-2.333.056c-.207-.197-.266-.487-.282-1.16l-.116-4.798a.75.75 0 1 0-1.5.036l.017.693a4.179 4.179 0 1 0 .532 5.94c.087.13.19.257.315.375a3.15 3.15 0 0 0 4.453-.107c.37-.388.557-.975.658-1.43a6.13 6.13 0 0 0 .145-1.388a.777.777 0 0 0-.006-.082a.763.763 0 0 0 .004-.077a8.75 8.75 0 1 0-3.337 6.875a.75.75 0 0 0-.928-1.178A7.25 7.25 0 0 1 4.755 12Zm7.25-2.679a2.679 2.679 0 1 0 0 5.358a2.679 2.679 0 0 0 0-5.358Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}