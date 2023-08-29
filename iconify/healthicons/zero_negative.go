package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthicons0Negative0)"><path d="M18 18a4 4 0 0 1 4-4h4a4 4 0 0 1 1.25.198L18 29.381V18Zm2.146 15.546l9.815-16.11c.026.184.039.373.039.564v12a4 4 0 0 1-4 4h-4c-.67 0-1.3-.163-1.854-.454Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM22 10a8 8 0 0 0-8 8v12c0 2.048.772 3.92 2.038 5.334A7.985 7.985 0 0 0 22 38h4a8 8 0 0 0 8-8V18c0-2.37-1.033-4.5-2.666-5.963A7.977 7.977 0 0 0 26 10h-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons0Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}