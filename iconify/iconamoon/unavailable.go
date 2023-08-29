package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unavailable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20a8 8 0 0 1-8-8H2c0 5.523 4.477 10 10 10v-2Zm0-16a8 8 0 0 1 8 8h2c0-5.523-4.477-10-10-10v2Zm-8 8a7.97 7.97 0 0 1 2.343-5.657L4.93 4.93A9.972 9.972 0 0 0 2 11.999h2Zm2.343-5.657A7.972 7.972 0 0 1 12 4V2a9.972 9.972 0 0 0-7.071 2.929l1.414 1.414Zm-1.414 0l12.728 12.728l1.414-1.414L6.343 4.929L4.93 6.343ZM20 12a7.97 7.97 0 0 1-2.343 5.657l1.414 1.414A9.972 9.972 0 0 0 22 12h-2Zm-2.343 5.657A7.972 7.972 0 0 1 12 20v2a9.972 9.972 0 0 0 7.071-2.929l-1.414-1.414Z"/>`),
		g.Group(children),
	)
}