package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsQNegative0)"><path d="m29.414 26.586l2.93 2.929A9.954 9.954 0 0 0 34 24c0-5.523-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10c2.038 0 3.934-.61 5.515-1.657l-2.93-2.929a2 2 0 1 1 2.83-2.828Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 10c-7.732 0-14 6.268-14 14s6.268 14 14 14c3.145 0 6.047-1.037 8.384-2.787l2.202 2.201a2 2 0 1 0 2.828-2.828l-2.201-2.202A13.938 13.938 0 0 0 38 24c0-7.732-6.268-14-14-14Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsQNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}