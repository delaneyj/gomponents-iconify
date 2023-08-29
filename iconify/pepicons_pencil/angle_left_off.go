package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="m12.384 15.68l-5-6l-.768.64l5 6l.768-.64Z"/><path d="m11.616 16.32l-5-6c-.427-.512.341-1.152.768-.64l5 6c.427.512-.341 1.152-.768.64Z"/><path d="m11.616 3.68l-5 6l.768.64l5-6l-.768-.64Z"/><path d="m12.384 4.32l-5 6c-.427.512-1.195-.128-.768-.64l5-6c.427-.512 1.195.128.768.64Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}