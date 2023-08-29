package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeCardLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.166" height="25.627" y="11.187" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.756" ry="1.756"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.345 36.813v3.931c0 .97.786 1.756 1.756 1.756h24.7c.97 0 1.756-.786 1.756-1.756V7.256c0-.97-.786-1.756-1.756-1.756H10.1c-.97 0-1.756.786-1.756 1.756v3.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.732 11.187h-3.11c-.97 0-1.756.786-1.756 1.756v22.114c0 .97.786 1.756 1.756 1.756h3.11m5.851-25.626h-3.11c-.97 0-1.756.786-1.756 1.756v22.114c0 .97.786 1.756 1.756 1.756h3.11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.939 18.129h-8.307a2.345 2.345 0 0 0-2.345 2.344v.61m10.652 8.862h-8.307a2.345 2.345 0 0 1-2.345-2.345v-1.218a2.345 2.345 0 0 1 2.345-2.345h8.307"/>`),
		g.Group(children),
	)
}