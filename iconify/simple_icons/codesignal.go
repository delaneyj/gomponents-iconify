package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codesignal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 1.212L13.012 2.787L12 5.62l-1.01-2.833L0 1.212L3.672 11.45l4.512.646l3.815 10.691l3.816-10.691l4.512-.646zm-3.625 4.406l-4.52.648l-.73 2.044l4.517-.647l-.734 2.047l-4.514.647L12 17.064l-2.393-6.707l-4.514-.647l-.735-2.047l4.518.647l-.73-2.044l-4.52-.648l-.735-2.047l6.676.956L12 11.345l2.434-6.818l6.676-.956Z"/>`),
		g.Group(children),
	)
}