package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M108 699V134c23-14 35-36 35-62c0-39-32-72-71-72C32 0 0 33 0 72c0 26 14 48 37 62v565c0 10 7 18 16 18h37c10 0 18-8 18-18zm35-533v324c0 19 10 27 22 18c23-17 43-30 61-39c37-17 63-25 87-26c30 0 53 6 76 17c22 10 42 26 65 41c13 8 29 13 47 14c30 2 72-5 121-45c11-9 20-32 20-49V97c0-19-8-26-20-17c-25 19-48 32-68 38c-41 11-73 9-100-7c-23-15-43-31-65-41c-23-11-46-17-76-17c-11 0-25 3-40 7c-27 8-62 25-108 57c-12 9-22 31-22 49z"/>`),
		g.Group(children),
	)
}