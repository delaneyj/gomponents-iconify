package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m356 538l156-155c11-11 12-29 3-42l151-127v466H0V14h554L400 226l-44-44c-12-12-33-12-45 0l-9 9c-25-40-70-66-120-66c-78 0-140 62-140 140c0 76 60 137 135 140l134 133c12 13 33 13 45 0zm-30-88l-86-68c24-12 44-32 57-56c15 20 26 34 27 35c1 2 4 4 7 4s6-1 9-5L592 14h74v160L337 451c-2 1-4 2-5 2c-4 0-6-3-6-3zM182 155c60 0 110 49 110 110c0 60-50 109-110 109S73 325 73 265c0-61 49-110 109-110z"/>`),
		g.Group(children),
	)
}