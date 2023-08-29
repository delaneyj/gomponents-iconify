package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fandango(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.664 6.956L8.05 8.496l1.14 4.224l5.615-1.54l1.145 4.22l-5.615 1.49l1.093 4.224l-5.615 1.49L4.42 17.54c.846-.995 1.194-2.386.846-3.728c-.398-1.342-1.392-2.385-2.584-2.832L1.29 5.763L12.57 2.78zm7.106-.198L18.932.05L0 5.068l1.838 6.758a3.046 3.046 0 0 1 2.385 2.236c.348 1.193-.1 2.385-.944 3.18l1.788 6.708L24 18.882l-1.79-6.708c-1.142-.2-2.086-1.043-2.434-2.236c-.298-1.193.1-2.435.994-3.18z"/>`),
		g.Group(children),
	)
}