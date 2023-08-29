package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1824 896q93 0 158.5 65.5T2048 1120v384q0 14-9 23t-23 9h-96v64q0 80-56 136t-136 56t-136-56t-56-136v-64H512v64q0 80-56 136t-136 56t-136-56t-56-136v-64H32q-14 0-23-9t-9-23v-384q0-93 65.5-158.5T224 896h28l105-419q23-94 104-157.5T640 256h128V32q0-14 9-23t23-9h448q14 0 23 9t9 23v224h128q98 0 179 63.5T1691 477l105 419h28zM320 1376q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47zm196-480h1016l-89-357q-2-8-14-17.5t-21-9.5H640q-9 0-21 9.5T605 539zm1212 480q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47z"/>`),
		g.Group(children),
	)
}