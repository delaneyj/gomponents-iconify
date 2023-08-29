package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wizard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.075 1005q-19 19-45 20t-44-17l-642-643q-18-18-17-44t20-45t45-20t44 17l642 643q18 18 17 44t-20 45zm-60-749l-80-46l-79 46l17-93l-66-65l90-12l38-86l39 86l90 12l-66 65zm-798-17l-129-130q-18-18-17-44t20-45t45-20t44 17l129 130q18 18 17 44t-20 45t-45 20t-44-17zm-25 231l39-86l38 86l90 12l-66 65l17 94l-79-46l-80 46l17-94l-65-65zm256 385l39-86l38 86l90 11l-66 66l17 93l-79-46l-79 46l17-93l-66-66z"/>`),
		g.Group(children),
	)
}