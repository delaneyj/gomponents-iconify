package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#akarIconsTumblrFill0)"><path fill="currentColor" d="M18.895 22.517c-.798.867-2.646 1.456-4.301 1.483h-.182c-5.557 0-6.766-4.164-6.766-6.594v-6.748H5.458a.454.454 0 0 1-.324-.137a.472.472 0 0 1-.134-.33V7.003a.81.81 0 0 1 .142-.458a.782.782 0 0 1 .376-.29c2.855-1.026 3.748-3.562 3.87-5.49c.035-.516.297-.765.738-.765H13.4a.451.451 0 0 1 .33.134a.468.468 0 0 1 .137.333V5.87h3.823c.121 0 .238.05.324.137a.472.472 0 0 1 .134.33v3.83a.472.472 0 0 1-.134.33a.454.454 0 0 1-.324.138h-3.84v6.245c0 1.568 1.015 2.001 1.64 2.001a4.537 4.537 0 0 0 1.488-.321a.973.973 0 0 1 .595-.106a.483.483 0 0 1 .34.37l1.012 3.014c.068.237.14.498-.03.68Z"/></g><defs><clipPath id="akarIconsTumblrFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}