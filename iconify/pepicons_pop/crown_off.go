package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m15.267 9.928l-.367.234a3 3 0 0 1-4.107-.867L10 8.105l-.793 1.19a3 3 0 0 1-4.107.867l-.367-.234L5.542 13h8.916l.809-3.072ZM3.654 6.87c-.768-.488-1.736.219-1.504 1.099l1.654 6.285A1 1 0 0 0 4.77 15h10.458a1 1 0 0 0 .967-.745l1.654-6.286c.232-.88-.736-1.587-1.504-1.099l-2.52 1.604a1 1 0 0 1-1.369-.289l-1.625-2.437a1 1 0 0 0-1.664 0L7.543 8.185a1 1 0 0 1-1.369.29L3.654 6.87Z" clip-rule="evenodd"/><path d="M11 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm8 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-16 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M4.5 16.25a1 1 0 0 1 1-1h9.251a1 1 0 1 1 0 2H5.5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}