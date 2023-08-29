package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFilledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 15.9L6.865 18l.533-5.13L4 9.04l4.965-1.07L11.5 3.5l2.535 4.469L19 9.039l-3.398 3.831l.533 5.13l-4.635-2.1Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M5.571 17.455L10 15.448l4.429 2.007a.5.5 0 0 0 .704-.507l-.51-4.91l3.251-3.668a.5.5 0 0 0-.269-.82L12.86 6.527l-2.425-4.274a.5.5 0 0 0-.87 0L7.14 6.527L2.395 7.55a.5.5 0 0 0-.27.82l3.253 3.667l-.51 4.911a.5.5 0 0 0 .703.507Zm4.223-3.01l-3.842 1.74l.443-4.263a.5.5 0 0 0-.123-.384l-2.83-3.191l4.128-.89a.5.5 0 0 0 .33-.242L10 3.513l2.1 3.702a.5.5 0 0 0 .33.242l4.128.89l-2.83 3.191a.5.5 0 0 0-.123.384l.443 4.263l-3.842-1.74a.5.5 0 0 0-.412 0Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}