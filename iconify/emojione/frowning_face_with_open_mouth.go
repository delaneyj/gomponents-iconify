package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrowningFaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><g fill="#664e27"><circle cx="20.5" cy="26" r="5"/><circle cx="43.5" cy="26" r="5"/><path d="M45.7 44c-1.5-3.6-5.1-6-13.7-6s-12.2 2.4-13.7 6c-.8 1.9.3 4 .3 4c.4 1.2 2.2 2 13.4 2c11.1 0 12.9-.8 13.4-2c0 0 1.1-2.1.3-4"/></g><path fill="#fff" d="M42 43c.1-.3 0-.6-.2-.8c0 0-2.2-2.2-9.8-2.2c-7.5 0-9.8 2.2-9.8 2.2c-.2.1-.2.5-.2.8l.2.6c.1.3.3.5.6.5h18.4c.2 0 .5-.2.6-.5l.2-.6"/>`),
		g.Group(children),
	)
}