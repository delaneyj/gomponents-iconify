package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopNutCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M18.156 19.992A2 2 0 0 1 16.42 21H9.58a2 2 0 0 1-1.736-1.008l-3.429-6a2 2 0 0 1 0-1.984l3.429-6A2 2 0 0 1 9.58 5h6.84a2 2 0 0 1 1.736 1.008l3.429 6a2 2 0 0 1 0 1.984l-3.429 6ZM9.58 19h6.84l3.428-6l-3.428-6H9.58l-3.428 6l3.428 6Z"/><path d="M13 11a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopNutCircleFilled0)"/></g>`),
		g.Group(children),
	)
}