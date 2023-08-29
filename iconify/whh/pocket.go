package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.56 896q-88 0-162-16t-140-54t-111.5-96t-72-146T.56 384V128q0-53 37.5-90.5T128.56 0h768q53 0 90.5 37.5t37.5 90.5v256q0 112-26.5 200t-72 146t-111.5 96t-140 54t-162 16zm295.5-617.5q-24.5-23.5-59.5-23.5t-59 23l-177 169l-178-169q-24-23-58.5-23t-59 23.5t-24.5 56.5t24 56l237 225q24 24 58.5 24t59.5-24l236-225q25-23 25-56t-24.5-56.5z"/>`),
		g.Group(children),
	)
}