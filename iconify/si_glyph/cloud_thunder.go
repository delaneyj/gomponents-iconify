package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudThunder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.125 6.915l-.998 2.046h6.891c-.139-1.433-1.061-2.543-2.189-2.543c-.25 0-.487.067-.713.168c-.694-1.041-1.83-1.722-3.118-1.722c-.207 0-.407.021-.606.056c-.682-1.135-1.869-1.889-3.223-1.889c-1.912 0-3.492 1.498-3.781 3.455c-.018 0-.033-.005-.051-.005c-1.281 0-2.318 1.11-2.318 2.479h4.663l.944-2.058l4.499.013z"/><path d="M9.701 10.071H8.59l.648-2.051H7.381l-1.029 2.925l1.336-.013l-.657 3.037l2.67-3.898z"/></g>`),
		g.Group(children),
	)
}