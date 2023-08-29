package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.02 4.971a9.941 9.941 0 1 0 0 14.05a9.941 9.941 0 0 0 0-14.05Zm-13.34.71a8.894 8.894 0 0 1 6.05-2.6a8.812 8.812 0 0 1-2.61 6.04a8.75 8.75 0 0 1-6.04 2.61a8.875 8.875 0 0 1 2.6-6.05Zm-2.58 7.05a9.772 9.772 0 0 0 6.73-2.9a9.8 9.8 0 0 0 2.9-6.73a8.908 8.908 0 0 1 5.23 2.24L5.34 17.951a8.881 8.881 0 0 1-2.24-5.22Zm8.18 8.17a8.872 8.872 0 0 1-5.23-2.24l12.61-12.62a8.91 8.91 0 0 1 2.24 5.24a9.86 9.86 0 0 0-9.62 9.62Zm7.04-2.59a8.856 8.856 0 0 1-6.04 2.61a8.851 8.851 0 0 1 8.64-8.64a8.847 8.847 0 0 1-2.6 6.03Z"/>`),
		g.Group(children),
	)
}