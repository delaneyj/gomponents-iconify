package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximumTempOneHundredTenTwoHundredThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M363 195q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216h106q10 0 16-6t6-15zm-86 106q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5z"/>`),
		g.Group(children),
	)
}