package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-3.205 10.519c-.185.382-.373.402-.291 0C12.715 10.48 12.572 8.248 11 8v4.75c0 .973-.448 1.82-1.639 2.203c-1.156.369-2.449-.016-2.752-.846c-.303-.83.377-1.84 1.518-2.256c.637-.232 1.375-.292 1.873-.101V5h1c0 2.355 4.065 1.839 1.795 6.519z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}