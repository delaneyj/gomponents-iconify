package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.306 7.758l.343 3.088l-.694.055a2.111 2.111 0 0 0-1.915 1.764a20.641 20.641 0 0 0 0 6.67A2.111 2.111 0 0 0 6.955 21.1l1.496.12c2.362.188 4.736.188 7.098 0l1.496-.12a2.111 2.111 0 0 0 1.915-1.764a20.64 20.64 0 0 0 0-6.67a2.111 2.111 0 0 0-1.915-1.764l-.694-.055l.343-3.088a5.01 5.01 0 0 0 0-1.095l-.023-.205a4.7 4.7 0 0 0-9.342 0l-.023.205a4.96 4.96 0 0 0 0 1.095ZM12.374 3.8A3.2 3.2 0 0 0 8.82 6.624l-.023.205a3.46 3.46 0 0 0 0 .764l.349 3.139c1.9-.122 3.807-.122 5.708 0l.349-3.14a3.46 3.46 0 0 0 0-.763l-.023-.205a3.2 3.2 0 0 0-2.806-2.825ZM12 14.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}