package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilAngleLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="m15.384 18.68l-5-6l-.768.64l5 6l.768-.64Z"/><path d="m14.616 19.32l-5-6c-.427-.512.341-1.152.768-.64l5 6c.427.512-.341 1.152-.768.64Z"/><path d="m14.616 6.68l-5 6l.768.64l5-6l-.768-.64Z"/><path d="m15.384 7.32l-5 6c-.427.512-1.195-.128-.768-.64l5-6c.427-.512 1.195.128.768.64Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilAngleLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}