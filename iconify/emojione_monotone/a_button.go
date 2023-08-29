package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M28.271 30h7.459l-3.729-9.938z"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.476 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.524-4.477-10-10-10zM41.733 46l-4.202-11.199h-11.06L22.267 46H17l12.01-32h5.982L47 46h-5.267z"/>`),
		g.Group(children),
	)
}