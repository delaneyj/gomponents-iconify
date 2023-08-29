package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiObjectsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.65 0-1.175-.313T10 20.85q-.825 0-1.413-.588T8 18.85V15.3q-1.475-.975-2.363-2.575T4.75 9.25q0-3.025 2.113-5.138T12 2q3.025 0 5.138 2.113T19.25 9.25q0 1.925-.888 3.5T16 15.3v3.55q0 .825-.588 1.413T14 20.85q-.3.525-.825.838T12 22Zm-2-3.15h4v-.9h-4v.9Zm0-1.9h4V16h-4v.95ZM9.8 14h1.45v-2.7l-1.7-1.7q-.2-.2-.2-.5t.225-.525q.225-.225.525-.225t.525.225L12 9.95l1.4-1.4q.2-.2.5-.2t.525.225q.225.225.225.525t-.225.525L12.75 11.3V14h1.45q1.35-.65 2.2-1.912t.85-2.838q0-2.2-1.525-3.725T12 4Q9.8 4 8.275 5.525T6.75 9.25q0 1.575.85 2.838T9.8 14ZM12 9.95ZM12 9Z"/>`),
		g.Group(children),
	)
}