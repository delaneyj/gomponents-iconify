package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestartCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.727 6a.75.75 0 0 0-1.5 0v1.023a5.665 5.665 0 0 0-7.299.715c-2.237 2.29-2.237 5.997 0 8.287a5.669 5.669 0 0 0 8.144 0a5.934 5.934 0 0 0 1.634-4.874a.75.75 0 0 0-1.49.182c.16 1.3-.25 2.653-1.217 3.644a4.169 4.169 0 0 1-5.998 0c-1.668-1.708-1.668-4.483 0-6.19a4.166 4.166 0 0 1 4.883-.822h-.558a.75.75 0 0 0 0 1.5h2.651a.75.75 0 0 0 .75-.75V6Z"/><path fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25ZM2.75 12a9.25 9.25 0 1 1 18.5 0a9.25 9.25 0 0 1-18.5 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}