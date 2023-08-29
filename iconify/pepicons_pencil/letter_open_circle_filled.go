package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLetterOpenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.5 11a.5.5 0 0 1 .5.5V20h14v-8.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5Z"/><path d="M6 8.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 .5.5v4.67a.5.5 0 0 1-.223.416l-6.5 4.33a.5.5 0 0 1-.554 0l-6.5-4.33A.5.5 0 0 1 6 13.17V8.5ZM7 9v3.902l6 3.997l6-3.997V9H7Z"/><path d="M12.723 5.084a.5.5 0 0 1 .554 0l4.5 3l-.554.832L13 6.101L8.777 8.916l-.554-.832l4.5-3Zm7.131 5.062l1 1l-.708.708l-1-1l.707-.708Zm-13 .708l-1 1l-.708-.708l1-1l.708.708ZM9.75 11a.25.25 0 0 1 .25-.25h6a.25.25 0 1 1 0 .5h-6a.25.25 0 0 1-.25-.25Zm.5 2a.25.25 0 0 1 .25-.25h5a.25.25 0 1 1 0 .5h-5a.25.25 0 0 1-.25-.25Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLetterOpenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}