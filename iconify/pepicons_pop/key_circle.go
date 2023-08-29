package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.718 15.307A1 1 0 0 0 16 14.61V12a5 5 0 1 0-6 0v8a3 3 0 1 0 6 0a1 1 0 0 0-.365-.772l-.226-.186l.088-.05a1 1 0 0 0 .503-.868v-.775a1 1 0 0 0-.292-.706l-.65-.653l.66-.683Zm-1.75-3.558A.998.998 0 0 1 14 12v2.206l-1.063 1.098a1 1 0 0 0 .01 1.402l.915.917l-.703.402a1 1 0 0 0-.14 1.64l.897.738A1 1 0 0 1 12 20v-8c0-.086.01-.17.031-.25a1 1 0 0 0-.483-1.124a3 3 0 1 1 2.904 0a1 1 0 0 0-.483 1.123Z"/><path d="M13 8.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}