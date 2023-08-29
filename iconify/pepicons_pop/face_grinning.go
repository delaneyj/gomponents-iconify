package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrinning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 17.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0 2a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z" clip-rule="evenodd"/><path d="M8.5 7.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M5.409 11.538A.75.75 0 0 1 6 11.25h8a.75.75 0 0 1 .728.932L14 12l.728.182l-.001.003l-.001.004l-.004.013a3.765 3.765 0 0 1-.053.188a8.544 8.544 0 0 1-.908 2.008c-.68 1.087-1.878 2.352-3.761 2.352c-1.883 0-3.081-1.265-3.761-2.352a8.544 8.544 0 0 1-.962-2.196l-.003-.013v-.004l-.001-.002L6 12l-.728.182a.75.75 0 0 1 .137-.644Zm1.654 1.212c.117.265.265.56.448.852c.57.913 1.372 1.648 2.489 1.648c1.117 0 1.919-.735 2.489-1.648c.183-.292.33-.587.448-.852H7.063Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}