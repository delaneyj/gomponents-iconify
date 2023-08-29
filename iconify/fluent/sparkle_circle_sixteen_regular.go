package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparkleCircleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.902 8.3a.42.42 0 0 0-.804 0l-.108.356a.5.5 0 0 1-.335.335l-.356.107a.42.42 0 0 0 0 .804l.356.107a.5.5 0 0 1 .335.335l.108.357c.12.399.685.399.804 0l.107-.356a.5.5 0 0 1 .335-.336l.357-.107a.42.42 0 0 0 0-.804l-.357-.107a.5.5 0 0 1-.335-.336L9.902 8.3ZM7.55 5.388a.583.583 0 0 0-1.1 0l-.17.483a.667.667 0 0 1-.407.407l-.484.171a.583.583 0 0 0 0 1.1l.484.171c.19.067.34.217.406.406l.171.484a.583.583 0 0 0 1.1 0l.172-.484a.667.667 0 0 1 .406-.406l.483-.171a.583.583 0 0 0 0-1.1l-.483-.171a.667.667 0 0 1-.406-.406l-.172-.484ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm6-5a5 5 0 1 0 0 10A5 5 0 0 0 8 3Z"/>`),
		g.Group(children),
	)
}