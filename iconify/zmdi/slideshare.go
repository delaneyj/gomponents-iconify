package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M153 132q21 0 35.5 14t14.5 33.5t-14.5 33T153 226t-36-13.5t-15-33t15-33.5t36-14zm119 0q21 0 35.5 13.5t14.5 33t-14.5 33.5t-35.5 14t-36-14t-15-33.5t15-33t36-13.5zm131 74q10-7 15-.5t-1 15.5q-29 36-88 60q26 89-22 131q-32 27-64 14q-27-10-26-42q0 1-.5-24.5T216 306l-4-1l-7-2v81q1 36-32 44q-36 9-65-23q-40-43-16-124q-60-25-89-60q-6-9-1-15.5t14 .5l4 3V44q0-17 12.5-29T61 3h300q16 0 26 12t10 29v165q2 0 6-3zm-27 16V63q0-22-6.5-30.5T345 24H79q-20 0-26.5 8.5T46 63v160q23 14 51 19.5t46 4.5t34 0q15-1 22 6q1 0 1.5 1l.5 1q9 8 15 12q1-22 27-20q16-1 34 0t46-5t53-20z"/>`),
		g.Group(children),
	)
}