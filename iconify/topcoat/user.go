package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func User(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M20.474 40.996c.347 0 .705.01 1.051 0c6.773-.1 13.221-.861 14.389-1.973c.27-2.463.598-3.655-10.325-8.822c-1.554-.871-.731-4.246-.731-4.246c3.289-2.482 5.547-8.231 5.547-12.838c0-7.961-3.59-11.726-8.877-12.117h-1.052c-5.278.391-8.867 4.156-8.867 12.117c0 4.607 2.248 10.354 5.539 12.838c0 0 .82 3.375-.725 4.246c-10.934 5.167-10.606 6.36-10.335 8.822c1.167 1.111 7.613 1.873 14.386 1.973z"/>`),
		g.Group(children),
	)
}