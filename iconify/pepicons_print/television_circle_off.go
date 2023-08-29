package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="m15 8.424l1.914-2.393c.625-.78 1.796.157 1.172.938L16.46 9H19a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-7.5a3 3 0 0 1-3-3v-5a3 3 0 0 1 3-3h2.04l-1.626-2.031c-.624-.781.547-1.718 1.172-.938L15 8.424Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M6.25 12v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Zm3 7a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h7.5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-7.5Z" clip-rule="evenodd"/><path d="m12.86 8.688l2-2.5c.416-.52 1.197.104.78.624l-2 2.5c-.416.52-1.197-.104-.78-.624Z"/><path d="m11.86 9.312l-2-2.5c-.417-.52.364-1.145.78-.624l2 2.5c.417.52-.364 1.145-.78.624ZM4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}