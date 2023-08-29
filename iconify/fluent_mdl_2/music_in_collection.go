package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicInCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1408q0 62-29 109t-76 80t-104 50t-111 17q-54 0-111-17t-103-49t-76-80t-30-110q0-61 29-109t76-80t104-50t111-17q51 0 100 12t92 39V226L768 450v1214q0 62-29 109t-76 80t-104 50t-111 17q-54 0-111-17t-103-49t-76-80t-30-110q0-61 29-109t76-80t104-50t111-17q51 0 100 12t92 39V350L1792 62v1346zM448 1792q27 0 60-8t62-23t50-40t20-57q0-33-20-57t-49-39t-63-24t-60-8q-27 0-60 8t-62 23t-50 40t-20 57q0 33 20 57t49 39t63 24t60 8zm1024-256q27 0 60-8t62-23t50-40t20-57q0-33-20-57t-49-39t-63-24t-60-8q-27 0-60 8t-62 23t-50 40t-20 57q0 33 20 57t49 39t63 24t60 8z"/>`),
		g.Group(children),
	)
}