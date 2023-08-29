package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.944 14.79a.5.5 0 0 1 .141-.277L17.163 4.435a.5.5 0 0 1 .707 0l3.89 3.89a.5.5 0 0 1 0 .706L11.68 19.11a.5.5 0 0 1-.277.14l-4.595.706a.5.5 0 0 1-.57-.57l.705-4.594Zm.964.314l-.577 3.76l3.759-.578l9.608-9.608l-3.181-3.182l-9.61 9.608Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m18.472 11.173l-3.537-3.53l.707-.708l3.536 3.53l-.706.708Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}