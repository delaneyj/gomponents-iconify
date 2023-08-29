package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q141 0 272 36t244 104t207 160t161 207t103 245t37 272q0 184-63 354t-183 311L359 246Q499 126 669 63t355-63zM0 1024q0-184 63-354t183-311l1443 1443q-140 120-310 183t-355 63q-141 0-272-36t-244-104t-207-160t-161-207t-103-245t-37-272z"/>`),
		g.Group(children),
	)
}