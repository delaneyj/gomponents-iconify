package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M7.132 3.646a.5.5 0 0 1 .707 0l3.03 3.03a.5.5 0 0 1 0 .708l-4.242 4.243a1.214 1.214 0 1 0 1.717 1.717L12.587 9.1a.5.5 0 0 1 .707 0l3.03 3.03a.5.5 0 0 1 0 .708l-4.243 4.242A6.5 6.5 0 0 1 2.89 7.89l4.243-4.243Zm.353 1.061l-3.889 3.89a5.5 5.5 0 1 0 7.778 7.777l3.89-3.889l-2.324-2.323l-3.889 3.889A2.214 2.214 0 1 1 5.92 10.92l3.889-3.89l-2.324-2.323Z"/><path d="m7.485 10.06l-3.03-3.03l.707-.707l3.03 3.03l-.707.708Zm5.455 5.456l-3.03-3.03l.707-.708l3.03 3.03l-.707.708Zm-.353-8.552a.5.5 0 0 1-.279-.65l1.175-2.937l1.53.51l1.01-2.103a.5.5 0 0 1 .9.432l-1.39 2.898l-1.471-.49l-.825 2.062a.5.5 0 0 1-.65.278Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}