package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AiHealing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4h2v1h-2z"/><circle cx="9" cy="4" r="1" fill="currentColor"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M18 10h-5V8l3-.01A3 3 0 0 0 19 5V3a3.009 3.009 0 0 0-3-3H8a3.009 3.009 0 0 0-3 3v2a3.009 3.009 0 0 0 3 3h3v2H6a3.009 3.009 0 0 0-3 3v8a3.009 3.009 0 0 0 3 3h12a3.009 3.009 0 0 0 3-3v-8a3.009 3.009 0 0 0-3-3ZM8 6a1.003 1.003 0 0 1-1-1V3a1.003 1.003 0 0 1 1-1h7.99A1.012 1.012 0 0 1 17 3v2a1.012 1.012 0 0 1-1.01 1Zm4.435 14.072L12 20.5l-.435-.432C10.02 18.541 9 17.534 9 16.298a1.712 1.712 0 0 1 1.65-1.798a1.74 1.74 0 0 1 1.35.683a1.74 1.74 0 0 1 1.35-.683A1.712 1.712 0 0 1 15 16.298c0 1.236-1.02 2.243-2.565 3.774Z"/>`),
		g.Group(children),
	)
}