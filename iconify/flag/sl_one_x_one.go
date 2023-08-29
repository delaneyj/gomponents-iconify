package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagSl1x10"><rect width="384" height="512" rx="4.6" ry="7.6"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagSl1x10)" transform="scale(1.33333 1)"><path fill="#0000cd" d="M0 341.7h512V512H0z"/><path fill="#fff" d="M0 171.4h512v170.3H0z"/><path fill="#00cd00" d="M0 0h512v171.4H0z"/></g>`),
		g.Group(children),
	)
}