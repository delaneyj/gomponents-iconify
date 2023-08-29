package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcologyScienceTestTubeExperimentLabScienceChemistryTestTubeSolution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h9M10 .5v10a3 3 0 0 1-6 0V.5"/>`),
		g.Group(children),
	)
}