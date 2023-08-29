package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.069 11.636C2.803 7.194 6.884 4 12 4c5.116 0 9.197 3.194 10.931 7.636a1 1 0 0 1 0 .728C21.197 16.806 17.116 20 12 20c-5.116 0-9.197-3.194-10.931-7.636a1 1 0 0 1 0-.728ZM12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}