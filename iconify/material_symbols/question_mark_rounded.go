package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMarkRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.6 8.075q0-1.075-.712-1.725T12 5.7q-.725 0-1.313.313t-1.012.912q-.4.575-1.088.663T7.4 7.225q-.35-.325-.388-.8t.238-.9q.8-1.2 2.038-1.863T12 3q2.425 0 3.938 1.375t1.512 3.6q0 1.125-.475 2.025t-1.75 2.125q-.925.875-1.25 1.363T13.55 14.6q-.1.6-.513 1t-.987.4q-.575 0-.988-.388t-.412-.962q0-.975.425-1.788T12.5 11.15q1.275-1.125 1.688-1.738t.412-1.337ZM12 22q-.825 0-1.413-.588T10 20q0-.825.588-1.413T12 18q.825 0 1.413.588T14 20q0 .825-.588 1.413T12 22Z"/>`),
		g.Group(children),
	)
}