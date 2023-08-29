package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M145.5 200.502c0 4.417-3.579 7.998-8.01 7.998h-16.98c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998v145.004zm-65 0c0 4.417-3.579 7.998-8.01 7.998H55.51c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998v145.004zm129 0c0 4.417-3.579 7.998-8.01 7.998h-16.98c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998v145.004z"/>`),
		g.Group(children),
	)
}