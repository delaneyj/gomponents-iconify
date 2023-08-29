package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarespace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M78 174L216 35q26-26 63-26t63 26q15 16-1 32q-15 16-31 0l-2-3q-13-11-30.5-10.5T248 67L106 209q-16 16-32 0t0-32q1-1 4-3zm311 4q-16-16-32 0l-3 3l-138 139q-13 13-31.5 13T153 320q-16-16-32 0q-6 7-6 16t6 15q1 3 4 3q26 24 61.5 23.5T248 352l141-143q6-7 6-16t-6-15zm-217 91l-3 3q-16 16 0 32q14 16 32 0l141-142q13-14 31.5-14t31.5 14q13 13 13 31.5T405 225L271 361q13 13 31 13t32-13l102-104q26-26 26-63.5T436 130t-63-26t-62 26zm-19-12l141-143q7-6 7-15.5T294 83q-6-7-15.5-7T263 83l-3 3l-138 139q-13 13-31.5 13T59 225t-13-31.5T59 162L193 26q-13-13-31-13t-32 13L28 130Q2 156 2 193.5T28 257t63 26t62-26z"/>`),
		g.Group(children),
	)
}