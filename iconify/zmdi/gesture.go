package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M37 83L0 47q5-6 19-20Q45 0 77 0q18 0 35.5 15T130 61q0 20-6 34t-21 36q-29 43-40 75q-5 18-2.5 29.5T71 247q9 0 24-18q16-17 48-58q18-22 46-41t60-19q42 0 62.5 27.5T335 200h52v53h-52q-6 69-36.5 100T235 384q-28 0-48.5-19.5T166 318q0-33 30-69.5t85-45.5v-3q-1-8-2.5-12.5t-5-10.5t-11-9t-18.5-3q-18 0-39 20t-48 53q-16 19-23.5 28T114 284.5T91 297q-30 10-56-9q-29-22-29-64q0-14 6-32.5T27 156t16.5-30T59 101.5T67 89q18-28 7-32q-8-3-37 26zm199 249q14 0 27.5-18t17.5-57q-30 8-45.5 27T220 316q0 7 5 11.5t11 4.5z"/>`),
		g.Group(children),
	)
}