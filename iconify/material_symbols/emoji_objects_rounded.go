package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiObjectsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.65 0-1.175-.313T10 20.85q-.825 0-1.413-.588T8 18.85V15.3q-1.475-.975-2.363-2.575T4.75 9.25q0-3.025 2.113-5.138T12 2q3.025 0 5.138 2.113T19.25 9.25q0 1.925-.888 3.5T16 15.3v3.55q0 .825-.588 1.413T14 20.85q-.3.525-.825.838T12 22Zm-2-3.15h4v-.9h-4v.9Zm0-1.9h4V16h-4v.95ZM11.25 14h1.5v-2.7l1.675-1.675q.225-.225.225-.525t-.225-.525Q14.2 8.35 13.9 8.35t-.5.2L12 9.95l-1.375-1.375Q10.4 8.35 10.1 8.35t-.525.225Q9.35 8.8 9.35 9.1t.2.5l1.7 1.7V14Z"/>`),
		g.Group(children),
	)
}