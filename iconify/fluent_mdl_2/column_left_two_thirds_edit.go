package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnLeftTwoThirdsEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1664h699l-32 128H0V256h1280v898l-128 128V384H128v1280zM2048 256v540q-29-19-61-31t-67-19V384h-384v514l-128 128V256h640zm-200 640q42 0 78 15t64 42t42 63t16 78q0 39-15 76t-43 65l-717 719l-377 94l94-377l717-718q28-28 65-42t76-15zm51 249q21-21 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-692 694l-34 135l135-34l692-693z"/>`),
		g.Group(children),
	)
}