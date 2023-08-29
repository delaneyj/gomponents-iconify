package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpsoff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 576h-70q-20 120-106.5 207T576 890v70q0 26-18.5 45t-45 19t-45.5-19t-19-45v-70q-121-20-207.5-107T134 576H64q-26 0-45-18.5T0 512t19-45.5T64 448h70q20-120 106.5-207T448 134V64q0-26 19-45t45.5-19t45 19T576 64v70q121 20 207.5 107T890 448h70q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5zM512 256q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}