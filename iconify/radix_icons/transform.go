package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.85 1.75a.9.9 0 0 1 .9-.9h1.5a.9.9 0 0 1 .9.9v.3h6.7v-.3a.9.9 0 0 1 .9-.9h1.5a.9.9 0 0 1 .9.9v1.5a.9.9 0 0 1-.9.9h-.3v6.7h.3a.9.9 0 0 1 .9.9v1.5a.9.9 0 0 1-.9.9h-1.5a.9.9 0 0 1-.9-.9v-.3h-6.7v.3a.9.9 0 0 1-.9.9h-1.5a.9.9 0 0 1-.9-.9v-1.5a.9.9 0 0 1 .9-.9h.3v-6.7h-.3a.9.9 0 0 1-.9-.9v-1.5Zm2.1 2.4v6.7h.3a.9.9 0 0 1 .9.9v.3h6.7v-.3a.9.9 0 0 1 .9-.9h.3v-6.7h-.3a.9.9 0 0 1-.9-.9v-.3h-6.7v.3a.9.9 0 0 1-.9.9h-.3Zm-.6-2.4h-.6v1.5h1.5v-1.5h-.9ZM5.1 6a.9.9 0 0 1 .9-.9h1a.9.9 0 0 1 .9.9v1a.91.91 0 0 1-.006.106A.908.908 0 0 1 8 7.1h1a.9.9 0 0 1 .9.9v1a.9.9 0 0 1-.9.9H8a.9.9 0 0 1-.9-.9V8c0-.036.002-.071.006-.106A.91.91 0 0 1 7 7.9H6a.9.9 0 0 1-.9-.9V6Zm1 0H6v1h1V6h-.9ZM8 8h1v1H8V8Zm-5.35 3.75h-.9v1.5h1.5v-1.5h-.6Zm9.1-10h1.5v1.5h-1.5v-1.5Zm.9 10h-.9v1.5h1.5v-1.5h-.6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}