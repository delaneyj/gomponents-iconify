package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bellalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m352.37 1024l-64-64q0-54-46.5-152.5T128.37 643l-22-22q-89-89-103.5-213t47.5-228q-22-13-36-35t-14-49q0-39 28.5-67.5T96.37 0q27 0 49.5 14t34.5 37q104-62 228-47t212 103l22 22q66 66 165.5 112.5t152.5 46.5l64 64q0 129-56 253t-147.5 215.5T605.37 968t-253 56zm561-625q-31-29-104.5-2t-161.5 98q26 15 41.5 40.5t15.5 56.5q0 47-32.5 79.5t-79.5 32.5q-31 0-56.5-15.5t-40.5-41.5q-71 88-98 161.5t2 104.5q37 37 138-12t208-156t156-208t12-138z"/>`),
		g.Group(children),
	)
}