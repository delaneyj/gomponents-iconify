package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M8.571 20.455L13 18.448l4.429 2.007a.5.5 0 0 0 .704-.507l-.51-4.91l3.251-3.668a.5.5 0 0 0-.269-.82L15.86 9.527l-2.425-4.274a.5.5 0 0 0-.87 0L10.14 9.527L5.395 10.55a.5.5 0 0 0-.27.82l3.253 3.667l-.51 4.911a.5.5 0 0 0 .703.507Zm4.223-3.01l-3.842 1.74l.443-4.263a.5.5 0 0 0-.123-.384l-2.83-3.191l4.128-.89a.5.5 0 0 0 .33-.242L13 6.513l2.1 3.702a.5.5 0 0 0 .33.242l4.128.89l-2.83 3.191a.5.5 0 0 0-.123.384l.443 4.263l-3.842-1.74a.5.5 0 0 0-.412 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}