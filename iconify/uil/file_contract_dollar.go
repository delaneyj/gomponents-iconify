package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileContractDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.71 20.29L20 18.56v-.31a2.75 2.75 0 0 0-2.75-2.75h-.34l-1.44-1.44a.67.67 0 0 1 .28-.06H19a1 1 0 0 0 0-2h-1.5v-1a1 1 0 0 0-2 0v1a2.74 2.74 0 0 0-1.47.59l-1.32-1.33a1 1 0 0 0-1.42 1.42L13 14.44v.31a2.75 2.75 0 0 0 2.75 2.75h.34l1.44 1.44a.67.67 0 0 1-.28.06H14a1 1 0 0 0 0 2h1.5v1a1 1 0 0 0 2 0v-1a2.74 2.74 0 0 0 1.5-.62l1.32 1.33a1 1 0 0 0 1.42 0a1 1 0 0 0-.03-1.42ZM10 19H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h6v4a1 1 0 0 0 1 1h5a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-5-5a1.07 1.07 0 0 0-.28-.19h-.09a1.31 1.31 0 0 0-.28-.1H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2Zm3-14.59L14.59 6H13Z"/>`),
		g.Group(children),
	)
}