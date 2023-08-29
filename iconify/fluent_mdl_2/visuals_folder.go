package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualsFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M168 1664h856v128H0V384q0-27 10-50t27-40t41-28t50-10h352q45 0 77 9t58 24t46 31t40 31t44 23t55 10h736q26 0 49 10t41 27t28 41t10 50v256h256q26 0 49 10t41 27t28 41t10 49q0 30-14 58l-99 199h-143l128-256H552l-384 768zm-40-207l309-618q17-33 47-52t68-19h984V512H800q-45 0-77-9t-58-24t-46-31t-40-31t-44-23t-55-10H128v1073zm1792-49h128v640h-128v-640zm-256-128h128v768h-128v-768zm-256 256h128v512h-128v-512zm-256 256h128v256h-128v-256z"/>`),
		g.Group(children),
	)
}