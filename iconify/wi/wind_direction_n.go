package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 770q0-209 103.5-386.5T384 103T769 0q156 0 298 61t245 164t164 245.5t61 299.5q0 156-61 298.5t-164 245t-245 163t-298 60.5q-157 0-299.5-61T224 1312T60.5 1067.5T0 770zm169 0q0 243 177 422q177 177 423 177q162 0 300-80.5t219-218.5t81-300t-81-300.5t-219-219T769 170t-300 80.5t-219 219T169 770zm334 258l256-600q2-11 10-11t10 11l255 600q4 11-1.5 17t-16.5 0l-237-89q-10-4-20 0l-239 89q-10 6-16 0t-1-17z" fill="currentColor"/>`),
		g.Group(children),
	)
}