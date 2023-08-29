package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2V2Zm2 1l.997-.077A1 1 0 0 0 5 2v1Zm16 3l.994.11A1 1 0 0 0 21 5v1ZM5.23 6l-.996.077L5.23 6Zm13.109 9.119l.07.997l-.07-.997Zm-10.355.74l-.071-.998l.071.997ZM3 4h2V2H3v2Zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74l.142 1.995Zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344l1.988.221ZM4.003 3.077l.23 3l1.995-.154l-.23-3l-1.995.154Zm.23 3l.617 8.017l1.995-.154l-.617-8.017l-1.994.154ZM21 5H5.23v2H21V5Zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886l.143 1.995ZM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763l-.142-1.995Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`),
		g.Group(children),
	)
}