package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthicons7Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 10a2 2 0 1 0 0 4h11.922a73.401 73.401 0 0 0-4.599 7.369c-2.42 4.444-4.613 9.602-5.302 14.343a2 2 0 0 0 3.958.576c.59-4.06 2.531-8.734 4.858-13.006c2.316-4.254 4.918-7.933 6.678-9.977A2 2 0 0 0 32 10H16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons7Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}