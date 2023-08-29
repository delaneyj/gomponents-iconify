package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsManNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM30 18.82c.217.513.41 1.19.564 1.995C30.97 22.931 31 25.273 31 26a2 2 0 1 0 4 0c0-.773-.026-3.431-.508-5.94c-.238-1.236-.616-2.607-1.265-3.717C32.576 15.228 31.405 14 29.556 14H18.444c-1.849 0-3.02 1.228-3.671 2.343c-.649 1.11-1.027 2.48-1.265 3.717C13.026 22.569 13 25.227 13 26a2 2 0 1 0 4 0c0-.727.03-3.069.436-5.185c.155-.805.347-1.482.564-1.994V42a2 2 0 0 0 3.994.153l1-13A2.02 2.02 0 0 0 23 29h2c0 .051.002.102.006.153l1 13A2 2 0 0 0 30 42V18.82ZM28 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsManNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}