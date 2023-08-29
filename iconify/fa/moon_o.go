package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1262 1175q-54 9-110 9q-182 0-337-90T570 849t-90-337q0-192 104-357q-201 60-328.5 229T128 768q0 130 51 248.5t136.5 204t204 136.5t248.5 51q144 0 273.5-61.5T1262 1175zm203-85q-94 203-283.5 324.5T768 1536q-156 0-298-61t-245-164t-164-245T0 768q0-153 57.5-292.5t156-241.5T449 69.5T739 1q44-2 61 39q18 41-15 72q-86 78-131.5 181.5T608 512q0 148 73 273t198 198t273 73q118 0 228-51q41-18 72 13q14 14 17.5 34t-4.5 38z"/>`),
		g.Group(children),
	)
}