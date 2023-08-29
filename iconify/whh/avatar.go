package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avatar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 1024V0h1024v1024H0zM960 64H64v768h896V64zM448 547v-13q-67-22-97.5-90T320 279q0-74 53.5-112.5T512 128t138.5 38.5T704 279q0 217-128 256v12q113 10 183.5 53T832 714q0 10-8.5 18t-26 13.5t-36 9.5t-47 6.5t-50 4t-54.5 2t-50.5.5h-95l-50.5-.5l-54.5-2l-50-4l-47-6.5l-36-9.5l-26-13.5l-8.5-18q2-71 72.5-114T448 547z"/>`),
		g.Group(children),
	)
}