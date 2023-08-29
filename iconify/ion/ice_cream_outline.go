package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="m352 256l-96 224l-62-145"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M299.42 223.48C291.74 239.75 275.18 252 256 252c-13.1 0-27-5-33.63-9.76C216.27 237.87 208 240 208 250v62a24.07 24.07 0 0 1-24 24h0a24.07 24.07 0 0 1-24-24v-56h-2c-35.35 0-62-28.65-62-64a64 64 0 0 1 64-64h8v-8a88 88 0 0 1 176 0v8h8a64 64 0 0 1 0 128c-21.78 0-42-13-52.59-32.51Z"/>`),
		g.Group(children),
	)
}