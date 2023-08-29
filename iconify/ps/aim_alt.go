package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AimAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M218 13q96 5 172 63.5T462 216q-3 64-63.5 109T242 371q-67 1-115-19q-3 0-3-2q1-1 4 0q29 7 46 3t22.5-12.5t.5-17.5q-4-8-27-14t-56-16t-55-26l1-1q3-8 4-10q2-3 4-3h64q2 0 4 3l7 16q4 11 15 12q12 0 15-11q1-4 0-7v-2l-1-1l-66-148q-2-5-6-5q-2 0-3 1t-1 2t-1 2q-1 1-57 126l-3 6l-9-12Q4 205 2 171q-5-69 57-115.5T218 13zm4 253V127q0-16-17-16q-16 0-16 16v139q0 16 16 16q17 0 17-16zm196 0V127q0-16-17-16q-9 0-12 4q-1 1-5 7l-48 108l-48-108q-1-2-2.5-4t-1.5-3q-4-4-13-4q-17 0-17 16v139q0 16 17 16t17-16v-65l33 67q4 10 14 10q11 0 17-11l32-66v65q0 16 17 16t17-16zm-299-42H81q-2 0-2-3l19-44q0-1 1-1q0-1 1-1l2 1v1l19 44q1 3-2 3z"/>`),
		g.Group(children),
	)
}