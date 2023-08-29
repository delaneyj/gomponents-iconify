package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Broom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1984 0q26 0 45 19t19 45q0 26-19 45l-730 730q54 69 81 150t28 168q0 103-39 197t-112 168l-68 68l-227 454L4 1086l454-227l68-68q73-73 167-112t198-39q87 0 168 27t150 82l730-730q19-19 45-19zm-915 1543L505 979l-285 142l707 707l142-285zm83-98q61-60 94-130t34-158q0-81-30-151t-84-124t-123-83t-152-31q-87 0-157 33t-131 95l549 549z"/>`),
		g.Group(children),
	)
}