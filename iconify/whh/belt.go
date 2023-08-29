package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Belt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1005.35 870l-135 135q-19 19-45.5 19t-45.5-19l-259-259l-2 3q-20 19-47 19t-46-19l-214-214q-19-19-19-46t19-47l3-2l-195-195q-19-19-19-45.5t19-45.5l135-135q19-19 45.5-19t45.5 19l195 195l3-3q19-20 46-20t46 20l214 214q19 19 19 46t-19 46l-3 3l259 259q19 19 19 45.5t-19 45.5zm-332.5-166q13.5 0 22.5-9t9-22.5t-9-23t-22.5-9.5t-23 9.5t-9.5 23t9.5 22.5t23 9zm-448-511q-13.5 0-22.5 9t-9 22.5t9 23t22.5 9.5t23-9.5t9.5-23t-9.5-22.5t-23-9zm264.5 63l-92 93l105 105q10 10 10 24t-10 24t-24 10t-24-10l-105-105l-92 92l215 216l233-233zm311.5 512q-13.5 0-23 9.5t-9.5 23t9.5 22.5t23 9t22.5-9t9-22.5t-9-23t-22.5-9.5z"/>`),
		g.Group(children),
	)
}