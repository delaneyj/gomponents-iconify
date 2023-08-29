package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 312q-28-28-74-46t-86-18q-31 0-69 10t-71 22t-90 22T0 312q41-21 100-66.5t105-86t96-72T384 56q39 0 72.5 17t55.5 47q22-30 55.5-47T640 56q33 0 83 31.5t96 72t105 86t100 66.5q-65 0-122-10t-90-22t-71-22t-69-10q-40 0-86 18t-74 46zm0 78q28-35 72-56.5t88-21.5q63 0 122 6.5t99 16t71 19t45 15.5l15 7q-40 20-94.5 56.5t-94 64.5t-85 49.5T672 568q-118 0-160-31q-42 31-160 31q-33 0-78.5-21.5t-85-49.5t-94-64.5T0 376q5-3 15-7t44-15t72-19.5t98-15.5t123-7q44 0 88 21.5t72 56.5z"/>`),
		g.Group(children),
	)
}