package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M201.002 111c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 144 48 140.421 48 135.99v-16.98c0-4.424 3.588-8.01 7.998-8.01h145.004zm0 65c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 209 48 205.421 48 200.99v-16.98c0-4.424 3.588-8.01 7.998-8.01h145.004zm0-129c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 80 48 76.421 48 71.99V55.01c0-4.424 3.588-8.01 7.998-8.01h145.004z"/>`),
		g.Group(children),
	)
}