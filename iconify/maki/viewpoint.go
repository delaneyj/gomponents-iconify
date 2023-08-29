package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viewpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.02 8.425a2.386 2.386 0 0 0-.46.44l-4.55-3.5a7.998 7.998 0 0 1 1.51-1.51Zm6.46-4.56l-3.5 4.55a2.397 2.397 0 0 1 .45.45l4.56-3.5a7.945 7.945 0 0 0-1.51-1.5Zm-5.176 6.148a1.5 1.5 0 1 0 1.683 1.291a1.5 1.5 0 0 0-1.683-1.291ZM6.43 2.235a7.933 7.933 0 0 0-2.06.55l2.2 5.32a2.044 2.044 0 0 1 .61-.17Zm2.14.01l-.75 5.69a2.49 2.49 0 0 1 .61.16l2.2-5.3a7.213 7.213 0 0 0-2.06-.55Z"/>`),
		g.Group(children),
	)
}