package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1019.5 432q-6.5 11-19.5 14.5t-24-2.5L768 340v197l124 247q6 11 3 24t-14.5 19.5t-24.5 3t-20-14.5L736 616L636 816q-7 11-20 14.5t-24-3t-14.5-19.5t2.5-24l124-247V343L528 444q-8 4-16 4q-5 0-16-4L320 343v194l124 247q6 11 3 24t-14.5 19.5t-24.5 3t-20-14.5L288 616L188 816q-7 11-20 14.5t-24.5-3T129 808t3-24l124-247V340L48 444q-11 6-24 2.5T4.5 432t-3-24T16 388l255-128q8-4 17-4t16 4l208 119l208-119q8-4 16-4q10 0 17 4l255 128q11 7 14.5 20t-3 24zM736 192q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28zm-448 0q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}