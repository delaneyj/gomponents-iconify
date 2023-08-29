package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCrownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m18.267 12.928l-.367.234a3 3 0 0 1-4.107-.867L13 11.105l-.793 1.19a3 3 0 0 1-4.107.867l-.367-.234L8.542 16h8.916l.809-3.072ZM6.654 9.87c-.768-.488-1.736.219-1.504 1.099l1.654 6.286A1 1 0 0 0 7.77 18h10.458a1 1 0 0 0 .967-.745l1.654-6.286c.232-.88-.736-1.587-1.504-1.099l-2.52 1.604a1 1 0 0 1-1.369-.289l-1.625-2.437a1 1 0 0 0-1.664 0l-1.625 2.437a1 1 0 0 1-1.369.29L6.654 9.87Z" clip-rule="evenodd"/><path d="M14 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm8 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-16 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M7.5 19.25a1 1 0 0 1 1-1h9.251a1 1 0 1 1 0 2H8.5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCrownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}