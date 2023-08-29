package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTrain0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTrain1" fill="currentColor"><path id="feTrain2" d="M15 18.932c-.966.068-2.002.068-3.04.068c-1.014 0-2.021 0-2.96-.063V21H6v-1l.711-1.423C5.09 18.106 4 17.11 4 15V7c0-4 3.955-4 8-4s8 0 8 4v8c0 2.092-1.095 3.09-2.717 3.566L18 20v1h-3v-2.068ZM7 16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM6 6v6h12V6H6Z"/></g></g>`),
		g.Group(children),
	)
}