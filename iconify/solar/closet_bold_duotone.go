package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosetBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.025 15.75c.067 2.021.303 3.235 1.146 4.078a3.1 3.1 0 0 0 1.079.697V22a.75.75 0 0 0 1.5 0v-1.129C6.82 21 8.193 21 10 21h4c1.806 0 3.18 0 4.25-.129V22a.75.75 0 0 0 1.5 0v-1.475a3.1 3.1 0 0 0 1.078-.697c.808-.808 1.059-1.956 1.137-3.828l.009-.25c.002-.392.026-2.308.026-2.75v-3c0-3.771 0-5.657-1.172-6.828C19.657 2 17.771 2 14 2h-1.25v13.75H2.025ZM6.25 18a.75.75 0 0 1 .75-.75h2a.75.75 0 0 1 0 1.5H7a.75.75 0 0 1-.75-.75Zm8.75-.75a.75.75 0 0 0 0 1.5h2a.75.75 0 0 0 0-1.5h-2Zm0-10a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M10 2h1.25v5.25H2.026c.066-2.021.302-3.235 1.146-4.078C4.343 2 6.229 2 10 2ZM2.002 8.75C2 9.142 2 9.558 2 10v3c0 .442.024 2.358.025 2.75h9.225v-7H2.002Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}