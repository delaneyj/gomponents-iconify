package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBlockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 16.5c0 .806.159 1.575.447 2.277c.056.136-.041.29-.188.289a72.194 72.194 0 0 1-4.91-.184l-1.515-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 4.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402c.07.43.125.862.167 1.295c.017.167-.17.277-.315.193A6 6 0 0 0 10.5 16.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5Zm4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5Zm-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}