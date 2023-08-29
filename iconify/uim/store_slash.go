package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18.586v-7.143a3.947 3.947 0 0 1-5-.8a3.957 3.957 0 0 1-1.789 1.154Z" opacity=".25"/><path fill="currentColor" d="M22 23a.997.997 0 0 1-.707-.293l-20-20a1 1 0 0 1 1.414-1.414l20 20A1 1 0 0 1 22 23Z"/><path fill="currentColor" d="M12.586 14H10a1 1 0 0 0-1 1v7h6v-5.586Z" opacity=".7"/><path fill="currentColor" d="M10 14h2.586l-2.49-2.49A3.84 3.84 0 0 1 9 10.642a3.998 3.998 0 0 1-5 .82V21a1 1 0 0 0 1 1h4v-7a1 1 0 0 1 1-1zm5 2.414V22h4a.993.993 0 0 0 .93-.655z" opacity=".25"/><path fill="currentColor" d="M13.211 11.797A3.957 3.957 0 0 0 15 10.643A3.998 3.998 0 0 0 22 8a1.006 1.006 0 0 0-.071-.371l-2-5A1 1 0 0 0 19 2H5a1 1 0 0 0-.929.629l-.008.02zM3.255 4.669L2.071 7.63A1.006 1.006 0 0 0 2 8a3.998 3.998 0 0 0 7 2.643a3.84 3.84 0 0 0 1.095.866z" opacity=".5"/>`),
		g.Group(children),
	)
}