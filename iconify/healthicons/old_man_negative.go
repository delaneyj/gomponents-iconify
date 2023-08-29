package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OldManNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsOldManNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm26.5 8.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0ZM35 25a3 3 0 0 0-3 3v.343a1 1 0 1 0 2 0V28a1 1 0 1 1 2 0v14.946a1 1 0 1 0 2 0V28a3 3 0 0 0-3-3Zm-7-5.567V42a2 2 0 0 1-3.99.199l-1-10A2.012 2.012 0 0 1 23 32h-2c0 .066-.003.133-.01.199l-1 10A2 2 0 0 1 16 42V27.919c-1.679-.223-3.09-.898-4.136-1.925A6.214 6.214 0 0 1 10 21.481A6.336 6.336 0 0 1 11.92 17c1.29-1.259 3.129-2 5.335-2h7.32c4.973 0 7.944 1.722 9.62 3.759a8.443 8.443 0 0 1 1.494 2.695c.146.44.26.893.309 1.353v.015l.002.006v.003c.002.139.002.174 0 .178V23v.009c.09 1.096-.758 1.893-1.899 1.984c-1.134.091-2.132-.713-2.241-1.8v-.002a3.488 3.488 0 0 0-.13-.52a4.492 4.492 0 0 0-.795-1.43c-.503-.612-1.375-1.35-2.935-1.808Zm-12-.263c-.5.152-.864.389-1.123.641a2.413 2.413 0 0 0-.72 1.708c-.006.64.232 1.24.674 1.674c.257.252.631.496 1.169.648V19.17Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsOldManNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}