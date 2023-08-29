package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFilledCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" opacity=".2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 18.9L9.865 21l.533-5.13L7 12.04l4.965-1.07L14.5 6.5l2.535 4.469L22 12.039l-3.398 3.831l.533 5.13l-4.635-2.1Z" opacity=".2"/><path d="M8.571 20.455L13 18.448l4.429 2.007a.5.5 0 0 0 .704-.507l-.51-4.91l3.251-3.668a.5.5 0 0 0-.269-.82L15.86 9.527l-2.425-4.274a.5.5 0 0 0-.87 0L10.14 9.527L5.395 10.55a.5.5 0 0 0-.27.82l3.253 3.667l-.51 4.911a.5.5 0 0 0 .703.507Zm4.223-3.01l-3.842 1.74l.443-4.263a.5.5 0 0 0-.123-.384l-2.83-3.191l4.128-.89a.5.5 0 0 0 .33-.242L13 6.513l2.1 3.702a.5.5 0 0 0 .33.242l4.128.89l-2.83 3.191a.5.5 0 0 0-.123.384l.443 4.263l-3.842-1.74a.5.5 0 0 0-.412 0Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}