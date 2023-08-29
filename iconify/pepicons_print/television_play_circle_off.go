package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionPlayCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M8.5 17v-5a3 3 0 0 1 3-3H19a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-7.5a3 3 0 0 1-3-3Z"/><path d="M9.5 12v5a2 2 0 0 0 2 2H19a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2h-7.5a2 2 0 0 0-2 2Zm-1 0v5a3 3 0 0 0 3 3H19a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Zm7-3l2-2.5l-2 2.5Z"/><path d="M17.968 5.914a.75.75 0 0 1 .118 1.055l-2 2.5a.75.75 0 1 1-1.172-.938l2-2.5a.75.75 0 0 1 1.054-.117ZM14.5 9l-2-2.5l2 2.5Z"/><path d="M12.031 5.914a.75.75 0 0 1 1.055.117l2 2.5a.75.75 0 1 1-1.172.938l-2-2.5a.75.75 0 0 1 .117-1.055Z"/></g><path fill-rule="evenodd" d="M7.25 12v5a2 2 0 0 0 2 2h7.5a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2h-7.5a2 2 0 0 0-2 2Zm-1 0v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.5 14.5L12 16v-3l2.5 1.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 14.5a.5.5 0 0 1-.243.429l-2.5 1.5A.5.5 0 0 1 11.5 16v-3a.5.5 0 0 1 .757-.429l2.5 1.5A.5.5 0 0 1 15 14.5Zm-2.5-.617v1.234l1.028-.617l-1.028-.617Zm3.062-7.773a.5.5 0 0 1 .078.702l-2 2.5a.5.5 0 1 1-.78-.624l2-2.5a.5.5 0 0 1 .702-.078Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.938 6.11a.5.5 0 0 1 .702.078l2 2.5a.5.5 0 1 1-.78.624l-2-2.5a.5.5 0 0 1 .078-.702Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}