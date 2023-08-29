package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Visa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.31 768h-896q-26 0-45-18.5T.31 704V576h1024v128q0 27-18.5 45.5t-45.5 18.5zM.31 64q0-26 18.5-45t45.5-19h896q27 0 45.5 19t18.5 45v64H.31V64zm329 448l52-320h83l-52 320h-83zm529-320h67l67 320h-77l-19-32h-96l-18 32h-87l124-296l1-2.5l4-6l7-7l11-6l16-2.5zm38 224l-29-137l-35 137h64zm-338 96q-25 0-49-4.5t-34-8.5l-11-4l12-70q17 14 50 19.5t60.5-1t27.5-26.5q0-13-18-25.5t-40-21.5t-40-29t-18-47q0-29 15.5-51t39.5-32t45-14.5t40-4.5q17 0 36 3.5t29 6.5l10 3l-12 67q-26-16-71.5-14.5t-45.5 25.5q0 12 18.5 23.5t40.5 21t40 30.5t18 48q0 36-24.5 61.5t-55.5 35t-63 9.5zm-456-320q16 0 25.5 7.5t11.5 14.5l2 7l28 144l10 47l79-220h90l-133 320h-87l-72-278q-22-13-56-24v-18h102z"/>`),
		g.Group(children),
	)
}