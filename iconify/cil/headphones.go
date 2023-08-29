package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 216h-48v-32c0-79.4-64.6-144-144-144s-144 64.6-144 144v32H64a48.055 48.055 0 0 0-48 48v128a48.055 48.055 0 0 0 48 48h48a32.036 32.036 0 0 0 32-32V184a112 112 0 0 1 224 0v224a32.036 32.036 0 0 0 32 32h48a48.055 48.055 0 0 0 48-48V264a48.055 48.055 0 0 0-48-48ZM112 408H64a16.019 16.019 0 0 1-16-16V264a16.019 16.019 0 0 1 16-16h48v56h.008l.012 104Zm352-16a16.019 16.019 0 0 1-16 16h-48V248h48a16.019 16.019 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}