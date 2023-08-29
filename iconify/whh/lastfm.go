package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lastfm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 64q-40 0-68 28t-28 68q0 38 42.5 67t85.5 29q80 0 136 56t56 136t-56 136t-136 56q-49 0-91-26.5T672 544q-49-129-96-224q-22-43-46-83t-37-59l-13-18q-9-29-60-62.5T320 64q-106 0-181 75T64 320t75 181t181 75t181-75l45 45q-44 45-102.5 69.5T320 640q-87 0-160.5-43T43 480.5T0 320t43-160.5T159.5 43T320 0q65 0 132 39.5t92 88.5q46 60 96 160q20 41 44 97t38 91l14 36q40 64 96 64q53 0 90.5-37.5T960 448t-37.5-90.5T832 320q-69 0-130.5-48T640 160q0-66 47-113T800 0q69 0 114.5 39t45.5 89h-64q-11-26-39-45t-57-19z"/>`),
		g.Group(children),
	)
}