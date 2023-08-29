package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m70.2 337.4l104.4 104.4L441.5 175L337 70.5L70.2 337.4zM.6 499.8c-2.3 9.3 2.3 13.9 11.6 11.6L151.4 465L47 360.6L.6 499.8zM487.9 24.1c-46.3-46.4-92.8-11.6-92.8-11.6c-7.6 5.8-34.8 34.8-34.8 34.8l104.4 104.4s28.9-27.2 34.8-34.8c0 0 34.8-46.3-11.6-92.8z"/>`),
		g.Group(children),
	)
}