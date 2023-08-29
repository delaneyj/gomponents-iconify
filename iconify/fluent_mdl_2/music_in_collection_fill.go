package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicInCollectionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1408q0 62-29 109t-76 80t-104 50t-111 17q-54 0-111-17t-103-49t-76-80t-30-110q0-61 29-109t76-80t104-50t111-17q51 0 100 12t92 39V226L768 450v1214q0 62-29 109t-76 80t-104 50t-111 17q-54 0-111-17t-103-49t-76-80t-30-110q0-61 29-109t76-80t104-50t111-17q51 0 100 12t92 39V350L1792 62v1346z"/>`),
		g.Group(children),
	)
}