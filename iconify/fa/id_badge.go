package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 1258q0 64-37 107t-91 43H384q-54 0-91-43t-37-107t9-118t29.5-104t61-78.5T452 929q80 75 188 75t188-75q56 0 96.5 28.5t61 78.5t29.5 104t9 118zM870 739q0 94-67.5 160.5T640 966t-162.5-66.5T410 739t67.5-160.5T640 512t162.5 66.5T870 739zm282 893V256H128v1376q0 13 9.5 22.5t22.5 9.5h960q13 0 22.5-9.5t9.5-22.5zm128-1472v1472q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V160Q0 94 47 47T160 0h352v96q0 14 9 23t23 9h192q14 0 23-9t9-23V0h352q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}