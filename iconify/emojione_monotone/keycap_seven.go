package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.478 2 2 6.477 2 12v40c0 5.523 4.478 10 10 10h40c5.522 0 10-4.477 10-10V12c0-5.523-4.478-10-10-10zm5 43.666A8.333 8.333 0 0 1 48.667 54H15.333A8.333 8.333 0 0 1 7 45.666V12.334A8.333 8.333 0 0 1 15.333 4h33.334A8.333 8.333 0 0 1 57 12.334v33.332z"/><path fill="currentColor" d="M23 18.775V13h20v4.514c-1.651 1.727-3.33 4.205-5.036 7.436a44.967 44.967 0 0 0-3.903 10.303c-.895 3.637-1.335 6.887-1.321 9.748H27.1c.098-4.484.968-9.059 2.612-13.72c1.644-4.663 3.84-8.831 6.586-12.504H23z"/>`),
		g.Group(children),
	)
}