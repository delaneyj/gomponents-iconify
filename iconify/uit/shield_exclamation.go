package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5zm7.99-9.247a.5.5 0 0 0-.593-.387a9.106 9.106 0 0 1-7.11-1.454a.5.5 0 0 0-.574 0a9.107 9.107 0 0 1-7.11 1.454a.498.498 0 0 0-.603.49v8.018a9.131 9.131 0 0 0 3.799 7.407l3.91 2.804a.497.497 0 0 0 .582 0l3.91-2.804A9.131 9.131 0 0 0 20 11.874V3.855a.498.498 0 0 0-.01-.102zM19 11.874a8.129 8.129 0 0 1-3.38 6.595L12 21.063L8.38 18.47A8.13 8.13 0 0 1 5 11.874v-7.42a10.082 10.082 0 0 0 7-1.53a10.084 10.084 0 0 0 7 1.53v7.42zm-7 2.001a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25z"/>`),
		g.Group(children),
	)
}