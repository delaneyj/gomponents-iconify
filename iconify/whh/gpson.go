package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 576h-70q-20 120-106.5 207T576 890v70q0 26-18.5 45t-45.5 19t-45.5-19t-18.5-45v-70q-121-20-207.5-107T134 576H64q-26 0-45-18.5T0 512t18.5-45.5T64 448h70q20-120 106.5-207T448 134V64q0-26 19-45t45.5-19t45 19T576 64v70q121 20 207.5 107T890 448h70q26 0 45 18.5t19 45.5t-18.5 45.5T960 576zM512 256q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75zm.5 448q-79.5 0-136-56T320 512.5t56.5-136t136-56.5T648 376.5t56 136T648 648t-135.5 56z"/>`),
		g.Group(children),
	)
}