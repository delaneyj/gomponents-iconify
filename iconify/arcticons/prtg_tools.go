package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrtgTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.633 29.754c.011-.23.017-.461.017-.694c0-7.61-6.16-13.77-13.76-13.77c-7.61 0-13.77 6.16-13.77 13.77c0 3.03.98 5.84 2.64 8.11L7.7 39.23a19.398 19.398 0 0 1-3.57-11.24c0-10.78 8.74-19.53 19.53-19.53c10.78 0 19.53 8.75 19.53 19.53a19.55 19.55 0 0 1-.642 4.984m-5.159-18.868l-5.726 6.388m-21.579-6.54l3.082 3.702"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.364 22.075l16.991 9.589l-2.86 2.579l-14.131-12.168z"/><circle cx="36.5" cy="36.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.357 36.643l-2.642 2.642m3.688-6.546l-1.65 1.65l.604 2.254l2.254.604l1.65-1.65"/>`),
		g.Group(children),
	)
}