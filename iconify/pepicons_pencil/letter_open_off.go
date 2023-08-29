package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.5 8a.5.5 0 0 1 .5.5V17h14V8.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 5.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 .5.5v4.67a.5.5 0 0 1-.223.416l-6.5 4.33a.5.5 0 0 1-.554 0l-6.5-4.33A.5.5 0 0 1 3 10.17V5.5ZM4 6v3.902l6 3.997l6-3.997V6H4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.723 2.084a.5.5 0 0 1 .554 0l4.5 3l-.554.832L10 3.101L5.777 5.916l-.554-.832l4.5-3Zm7.131 5.062l1 1l-.708.708l-1-1l.707-.708Zm-13 .708l-1 1l-.708-.708l1-1l.708.708ZM6.75 8A.25.25 0 0 1 7 7.75h6a.25.25 0 1 1 0 .5H7A.25.25 0 0 1 6.75 8Zm.5 2a.25.25 0 0 1 .25-.25h5a.25.25 0 1 1 0 .5h-5a.25.25 0 0 1-.25-.25Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}