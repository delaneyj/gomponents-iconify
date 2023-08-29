package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrinningCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFaceGrinningCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 20.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0 2a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M8.409 14.538A.75.75 0 0 1 9 14.25h8a.75.75 0 0 1 .728.932L17 15l.728.182l-.001.003l-.001.004l-.003.013a3.765 3.765 0 0 1-.054.188a8.544 8.544 0 0 1-.908 2.008c-.68 1.087-1.878 2.352-3.761 2.352c-1.883 0-3.081-1.265-3.761-2.352a8.544 8.544 0 0 1-.962-2.196l-.003-.013v-.004l-.001-.002L9 15l-.728.182a.75.75 0 0 1 .137-.644Zm1.654 1.212c.117.265.265.56.448.852c.57.913 1.372 1.648 2.489 1.648c1.117 0 1.919-.735 2.489-1.648c.183-.292.33-.587.448-.852h-5.874Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFaceGrinningCircleFilled0)"/></g>`),
		g.Group(children),
	)
}