package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1472 1472q0 26-19 45t-45 19H946q1 17 6 87.5t5 108.5q0 25-18 42.5t-43 17.5H576q-25 0-43-17.5t-18-42.5q0-38 5-108.5t6-87.5H64q-26 0-45-19t-19-45t19-45l402-403H192q-26 0-45-19t-19-45t19-45l402-403H352q-26 0-45-19t-19-45t19-45L691 19q19-19 45-19t45 19l384 384q19 19 19 45t-19 45t-45 19H923l402 403q19 19 19 45t-19 45t-45 19h-229l402 403q19 19 19 45z"/>`),
		g.Group(children),
	)
}