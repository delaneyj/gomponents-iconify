package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecisionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M490 1343q-25 0-48-10t-42-28L37 942Q0 905 0 852q0-27 10-50t27-40t41-27t50-10q26 0 49 9t42 28l362 362q18 18 27 41t10 50q0 26-10 49t-27 41t-41 28t-50 10zm724-724q-25 0-48-10t-42-28L762 219q-38-38-38-91q0-27 10-50t27-40t41-28t50-10q53 0 90 37l363 363q18 18 27 41t10 50q0 26-10 49t-27 41t-41 28t-50 10zm-520 391L332 648l316-316l362 362l-316 316zm1244 838q19 19 19 45t-19 45t-45 19q-26 0-45-19L897 988l91-91l950 951zM0 1920h1024v128H0v-128zm128-256h768v128H128v-128z"/>`),
		g.Group(children),
	)
}