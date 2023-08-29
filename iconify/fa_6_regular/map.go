package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M565.6 36.2C572.1 40.7 576 48.1 576 56v336c0 10-6.2 18.9-15.5 22.4l-168 64c-5.2 2-10.9 2.1-16.1.3l-183.9-61.2l-160 61c-7.4 2.8-15.7 1.8-22.2-2.7S0 463.9 0 456V120c0-10 6.1-18.9 15.5-22.4l168-64c5.2-2 10.9-2.1 16.1-.3l183.9 61.2l160-61c7.4-2.8 15.7-1.8 22.2 2.7zM48 136.5v284.7l120-45.7V90.8L48 136.5zm312 286.2V137.3l-144-48v285.4l144 48zm48-1.5l120-45.7V90.8l-120 45.7v284.7z"/>`),
		g.Group(children),
	)
}