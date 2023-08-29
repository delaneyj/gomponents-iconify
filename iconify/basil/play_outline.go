package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.266 10.484a1.917 1.917 0 0 1 0 3.032a35.764 35.764 0 0 1-9.916 5.416l-.653.232c-1.248.443-2.566-.401-2.735-1.69a42.49 42.49 0 0 1 0-10.948c.169-1.289 1.487-2.133 2.735-1.69l.653.232a35.765 35.765 0 0 1 9.916 5.416Zm-.918 1.846a.417.417 0 0 0 0-.66a34.262 34.262 0 0 0-9.5-5.189l-.653-.232a.572.572 0 0 0-.746.472a40.99 40.99 0 0 0 0 10.558a.572.572 0 0 0 .746.471l.653-.231a34.263 34.263 0 0 0 9.5-5.19Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}