package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsRupeeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 10a2 2 0 1 0 0 4h2c1.48 0 2.773.804 3.465 2H18a2 2 0 1 0 0 4h5.465A3.998 3.998 0 0 1 20 22h-2a2 2 0 0 0-1.664 3.11l8 12a2 2 0 1 0 3.328-2.22l-6.037-9.056A8.012 8.012 0 0 0 27.748 20H30a2 2 0 1 0 0-4h-2.252a7.954 7.954 0 0 0-.818-2H30a2 2 0 1 0 0-4H18Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRupeeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}