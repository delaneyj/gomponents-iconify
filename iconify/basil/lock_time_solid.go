package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTimeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.306 6.758l.347 3.122l-.476.042a2.389 2.389 0 0 0-2.154 2.028a24.113 24.113 0 0 0 0 7.1a2.39 2.39 0 0 0 2.154 2.028l1.134.1c1.22.107 2.444.161 3.668.162c.175 0 .266-.212.154-.346a7.002 7.002 0 0 1 3.947-11.35a.416.416 0 0 0 .333-.357l.28-2.529a4.89 4.89 0 0 0 0-1.095l-.022-.205a4.7 4.7 0 0 0-9.342 0l-.023.205a4.96 4.96 0 0 0 0 1.095ZM10.374 2.8A3.2 3.2 0 0 0 6.82 5.624l-.023.205a3.46 3.46 0 0 0 0 .764l.352 3.164a42.166 42.166 0 0 1 5.702 0l.352-3.164a3.46 3.46 0 0 0 0-.764l-.023-.205a3.2 3.2 0 0 0-2.806-2.825Z" clip-rule="evenodd"/><path fill="currentColor" d="M16.25 15a.75.75 0 0 0-1.5 0v1.773c0 .24.115.465.309.606l1 .727a.75.75 0 1 0 .882-1.213l-.691-.502V15Z"/><path fill="currentColor" fill-rule="evenodd" d="M15.5 22a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}