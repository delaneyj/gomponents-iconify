package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewTeamProject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M608 256q45 0 77 9t58 24t46 31t40 31t44 23t55 10h992q27 0 50 10t40 27t28 41t10 50v896h-128V512H928q-31 0-54 9t-44 24t-41 31t-45 31t-58 23t-78 10H128v1024h1024v128H0V384q0-27 10-50t27-40t41-28t50-10h480zm0 256q24 0 42-4t33-13t29-20t32-27q-17-15-31-26t-30-20t-33-13t-42-5H128v128h480zm1184 1152h256v128h-256v256h-128v-256h-256v-128h256v-256h128v256zm-339-723l-429 430l-429-430l90-90l339 338l339-338l90 90z"/>`),
		g.Group(children),
	)
}