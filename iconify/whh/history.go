package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320.9 704v-32l131-177q5-20 22-33.5t39-13.5q21 0 38 12.5t23 32.5l195 243v32h-32l-225-181l-159 117h-32zm192 320q-104 0-197.5-39.5T150.9 874l91-90q53 53 123 82.5t148 29.5q104 0 192.5-51.5t140-140T896.9 512t-51.5-192.5t-140-140T512.9 128q-94 0-175 43t-135 116l54 54q0 18-12.5 30.5T214.9 384h-171q-18 0-30.5-12.5T.9 341V171q0-17 12.5-29.5T43.9 129l68 67q71-91 176-143.5T512.9 0q104 0 199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5z"/>`),
		g.Group(children),
	)
}