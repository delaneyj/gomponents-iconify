package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpthalmologyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsOpthalmologyNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm33 15s-8.25-9-14.96-9C17.328 15 9 24 9 24s8.329 9 15.04 9C30.75 33 39 24 39 24Zm-23.776-2.916A43.969 43.969 0 0 0 11.83 24a43.969 43.969 0 0 0 3.394 2.916c.516.397 1.053.79 1.605 1.166A8.213 8.213 0 0 1 15.75 24a8.21 8.21 0 0 1 1.079-4.082c-.552.376-1.09.769-1.605 1.166Zm15.915-1.222A8.212 8.212 0 0 1 32.25 24a8.212 8.212 0 0 1-1.111 4.138a33.522 33.522 0 0 0 1.671-1.218A43.399 43.399 0 0 0 36.183 24a43.377 43.377 0 0 0-3.373-2.92a33.522 33.522 0 0 0-1.671-1.218ZM30.25 24a6.25 6.25 0 1 1-12.5 0a6.25 6.25 0 0 1 12.5 0ZM24 27.75a3.75 3.75 0 0 0 2.849-6.189a1.125 1.125 0 1 1-.41-.41A3.75 3.75 0 1 0 24 27.75Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsOpthalmologyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}