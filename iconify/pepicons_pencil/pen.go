package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.944 11.79a.5.5 0 0 1 .141-.277L14.163 1.435a.5.5 0 0 1 .707 0l3.89 3.89a.5.5 0 0 1 0 .706L8.68 16.11a.5.5 0 0 1-.277.14l-4.595.706a.5.5 0 0 1-.57-.57l.705-4.594Zm.964.314l-.577 3.76l3.759-.578l9.608-9.608l-3.181-3.182l-9.61 9.608Z"/><path d="m15.472 8.173l-3.537-3.53l.707-.708l3.536 3.53l-.706.708Z"/></g>`),
		g.Group(children),
	)
}