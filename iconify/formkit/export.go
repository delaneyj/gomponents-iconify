package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 14c-.28 0-.5-.22-.5-.5v-8c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M10.5 9a.47.47 0 0 1-.35-.15L7.5 6.2L4.85 8.85c-.2.2-.51.2-.71 0s-.2-.51 0-.71l3-3c.2-.2.51-.2.71 0l3 3c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M10.07 12.94c-.19 0-.37-.11-.45-.29c-.11-.25 0-.55.25-.66c1.91-.87 3.14-2.74 3.14-4.77c0-2.91-2.47-5.28-5.5-5.28S2 4.31 2 7.22c0 2.02 1.23 3.9 3.14 4.77c.25.11.36.41.25.66c-.11.25-.41.36-.66.25c-2.26-1.03-3.72-3.26-3.72-5.68C1 3.75 3.92.94 7.5.94S14 3.76 14 7.22c0 2.41-1.46 4.64-3.72 5.68c-.07.03-.14.05-.21.05Z"/>`),
		g.Group(children),
	)
}