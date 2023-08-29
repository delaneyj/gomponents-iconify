package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.718 12.307A1 1 0 0 0 13 11.61V9a5 5 0 1 0-6 0v8a3 3 0 1 0 6 0a1 1 0 0 0-.365-.772l-.226-.186l.088-.05a1 1 0 0 0 .503-.868v-.775a1 1 0 0 0-.292-.706l-.65-.653l.66-.683Zm-1.75-3.558A.998.998 0 0 1 11 9v2.206l-1.063 1.098a1 1 0 0 0 .01 1.402l.915.917l-.703.402a1 1 0 0 0-.14 1.64l.897.738A1 1 0 0 1 9 17V9c0-.086.01-.17.031-.25a1 1 0 0 0-.483-1.124a3 3 0 1 1 2.904 0a1 1 0 0 0-.483 1.123Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}