package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M128 1664h1408V640H128v1024zM512 448V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23zm768 0V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23zm384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90z"/>`),
		g.Group(children),
	)
}