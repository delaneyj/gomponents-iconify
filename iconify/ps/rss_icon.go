package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M67 301q-28 0-46 18.5T3 365q0 28 18 46t46 18q27 0 45.5-18t18.5-46q0-27-18.5-45.5T67 301zm0 86q-22 0-22-22q0-9 6-15t16-6q9 0 15 6t6 15q0 10-6 16t-15 6zm0-235q-10 0-16 6t-6 15q0 22 22 22q71 0 120.5 49.5T237 365q0 22 22 22q9 0 15-6t6-16q0-88-62.5-150.5T67 152zM67 3Q45 3 45 24t22 21q133 0 226.5 93.5T387 365q0 22 21 22t21-22q0-150-106-256T67 3z"/>`),
		g.Group(children),
	)
}