package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crescentmoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FDE364" d="M503.851 210.204C492.563 120.657 434.38 44.485 355.192 7.235c-11.279-5.306-22.337 7.572-15.47 17.982c18.48 28.029 30.919 60.278 35.273 94.838c18.733 148.659-106.281 273.673-254.94 254.94c-34.56-4.354-66.81-16.793-94.839-35.273c-10.41-6.866-23.287 4.191-17.982 15.478c37.25 79.182 113.422 137.364 202.969 148.651c171.226 21.579 315.226-122.414 293.648-293.647"/>`),
		g.Group(children),
	)
}