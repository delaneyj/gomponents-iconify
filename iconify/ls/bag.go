package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M680 221v120c-95 35-213 55-340 55S95 376 0 341V221c0-22 18-40 40-40h165v-30c0-46 38-85 85-85h100c47 0 85 39 85 85v30h165c22 0 40 18 40 40zm-425-70v30h170v-30c0-19-16-35-35-35H290c-19 0-35 16-35 35zm35 183v20c0 6 4 10 10 10h80c6 0 10-4 10-10v-20c0-6-4-10-10-10h-80c-6 0-10 4-10 10zM0 602V384c96 33 213 52 340 52s244-19 340-52v218c0 22-18 40-40 40H40c-22 0-40-18-40-40zm390-103v-20c0-6-4-10-10-10h-80c-6 0-10 4-10 10v20c0 6 4 10 10 10h80c6 0 10-4 10-10z"/>`),
		g.Group(children),
	)
}