package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.011 32s-1.12-.203-2.021-.203c-.812 0-2.031.203-2.031.203l.255-13.593C12.323 15.146 6.782 5.214 3.323 0c1.74.396 2.469.369 4.219 0l.027.047c2.203 3.901 5.572 9.339 8.421 14.052C18.803 9.47 23.23 2.24 24.427 0c1.36.36 2.729.344 4.251 0c-1.599 2.156-7.423 12.229-10.937 18.407L18.001 32z"/>`),
		g.Group(children),
	)
}