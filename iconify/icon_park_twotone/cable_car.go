package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CableCar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCableCar0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="bevel" d="M13.28 42a2 2 0 0 1-1.816-1.162l-5.077-11a2 2 0 0 1 0-1.676l5.077-11A2 2 0 0 1 13.28 16h21.44a2 2 0 0 1 1.816 1.162l5.077 11a2 2 0 0 1 0 1.676l-5.077 11A2 2 0 0 1 34.72 42H13.28Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 23h9v10H8m31-10H25v9h15M38 4s-9.5 3.5-17.436 4.394C12.628 9.288 8 9 8 9m18-1v8l-9-7"/><path stroke-linecap="round" stroke-linejoin="round" d="m10.615 19l-4.228 9.162a2 2 0 0 0 0 1.676L9 35.5"/><path stroke-linecap="round" stroke-linejoin="bevel" d="m37.385 19l4.228 9.162a2 2 0 0 1 0 1.676L39 35.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCableCar0)"/>`),
		g.Group(children),
	)
}