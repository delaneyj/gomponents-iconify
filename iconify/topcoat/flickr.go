package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M27.96 13.65c4.06 0 7.36 3.289 7.36 7.35s-3.301 7.35-7.36 7.35s-7.35-3.289-7.35-7.35s3.29-7.35 7.35-7.35zm-15.92 0c4.06 0 7.35 3.289 7.35 7.35s-3.29 7.35-7.35 7.35S4.68 25.061 4.68 21s3.3-7.35 7.36-7.35zM.5 1.5v39h39v-39H.5z"/>`),
		g.Group(children),
	)
}