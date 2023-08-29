package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenFolderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 768q26 0 49 10t41 27t28 41t10 49q0 30-14 58l-419 839H0V384q0-27 10-50t27-40t41-28t50-10h352q45 0 77 9t58 24t46 31t40 31t44 23t55 10h736q27 0 50 10t40 27t28 41t10 50v256h256zM128 1457l309-618q17-33 47-52t68-19h984V512H800q-45 0-77-9t-58-24t-46-31t-40-31t-44-23t-55-10H128v1073zm1792-561H552l-384 768h1368l384-768z"/>`),
		g.Group(children),
	)
}