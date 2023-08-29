package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximumTempTwoHundredThreeHundredNinety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 429h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216h106q10 0 16-6t6-15q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21zm171-128q0 9-6.5 15.5T171 323q-9 0-15.5-6.5T149 301q0-8 6.5-14.5T171 280q8 0 14.5 6.5T192 301zm85 0q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5zm86 0q0 9-6.5 15.5T341 323q-8 0-14.5-6.5T320 301q0-8 6.5-14.5T341 280q9 0 15.5 6.5T363 301z"/>`),
		g.Group(children),
	)
}