package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.557 12.763a.5.5 0 0 1 .141-.277l9.394-9.394a.5.5 0 0 1 .707 0l3.625 3.625a.5.5 0 0 1 0 .707l-9.394 9.394a.5.5 0 0 1-.277.14l-4.283.658a.5.5 0 0 1-.57-.57l.657-4.283Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M3.944 11.79a.5.5 0 0 1 .141-.277L14.163 1.435a.5.5 0 0 1 .707 0l3.89 3.89a.5.5 0 0 1 0 .706L8.68 16.11a.5.5 0 0 1-.277.14l-4.595.706a.5.5 0 0 1-.57-.57l.705-4.594Zm.964.314l-.577 3.76l3.759-.578l9.608-9.608l-3.181-3.182l-9.61 9.608Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m15.472 8.173l-3.537-3.53l.707-.708l3.536 3.53l-.706.708Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}