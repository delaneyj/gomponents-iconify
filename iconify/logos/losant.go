package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Losant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FF495C" d="m181.929 35.112l.372-.747l-.123-33.428l-.249-.937L72.758 111.177l-.7.926v32.1l.628 1.354z"/><path fill="#284FFF" d="M181.908 0v35.133l38.784 38.016l-37.632 38.009L72.576.485L0 73.149l72.697 72.394l.034-34.392l-37.996-38.009l38.023-38.016l110.679 110.679l72.314-72.906z"/>`),
		g.Group(children),
	)
}