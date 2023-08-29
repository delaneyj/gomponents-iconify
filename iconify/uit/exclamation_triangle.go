package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15.875a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25zM12 14a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5zm10.713 3.547L14.564 3.479a2.964 2.964 0 0 0-5.128 0l-8.15 14.07A2.965 2.965 0 0 0 3.852 22h16.294a2.966 2.966 0 0 0 2.567-4.453zM20.146 21H3.852a1.965 1.965 0 0 1-1.7-2.95L10.3 3.98a1.943 1.943 0 0 1 1.699-.979a1.944 1.944 0 0 1 1.7.98l8.148 14.068A1.966 1.966 0 0 1 20.146 21z"/>`),
		g.Group(children),
	)
}