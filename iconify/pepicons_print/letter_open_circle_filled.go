package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M5.5 11a.5.5 0 0 1 .5.5V20h14v-8.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5Z"/><path d="M6 8.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 .5.5v4.67a.5.5 0 0 1-.223.416l-6.5 4.33a.5.5 0 0 1-.554 0l-6.5-4.33A.5.5 0 0 1 6 13.17V8.5ZM7 9v3.902l6 3.997l6-3.997V9H7Z"/><path d="M12.723 5.084a.5.5 0 0 1 .554 0l4.5 3l-.554.832L13 6.101L8.777 8.916l-.554-.832l4.5-3Zm7.131 5.062l1 1l-.708.708l-1-1l.707-.708Zm-13 .708l-1 1l-.708-.708l1-1l.708.708ZM9.75 11a.25.25 0 0 1 .25-.25h6a.25.25 0 1 1 0 .5h-6a.25.25 0 0 1-.25-.25Zm.5 2a.25.25 0 0 1 .25-.25h5a.25.25 0 1 1 0 .5h-5a.25.25 0 0 1-.25-.25Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}