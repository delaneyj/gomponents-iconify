package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.229 11.229c-1.583 1.582-3.417 3.096-4.142 2.371c-1.037-1.037-1.677-1.941-3.965-.102c-2.287 1.838-.53 3.064.475 4.068c1.16 1.16 5.484.062 9.758-4.211c4.273-4.274 5.368-8.598 4.207-9.758c-1.005-1.006-2.225-2.762-4.063-.475c-1.839 2.287-.936 2.927.103 3.965c.722.725-.791 2.559-2.373 4.142z"/>`),
		g.Group(children),
	)
}