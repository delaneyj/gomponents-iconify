package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16.5v.5c1.691-2.578 3.6-3.953 6-4v3c0 .551.511 1 1.143 1c.364 0 .675-.158.883-.391C17.959 14.58 22 10.5 22 10.5s-4.041-4.082-5.975-6.137A1.262 1.262 0 0 0 15.143 4C14.511 4 14 4.447 14 5v3c-4.66 0-6 4.871-6 8.5zM5 21h14a1 1 0 0 0 1-1v-6.046c-.664.676-1.364 1.393-2 2.047V19H6V7h7V5H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1z"/>`),
		g.Group(children),
	)
}