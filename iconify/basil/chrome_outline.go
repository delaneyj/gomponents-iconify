package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.11 6.27A9.706 9.706 0 0 0 2.25 12c0 4.997 3.759 9.116 8.604 9.683A9.844 9.844 0 0 0 12 21.75A9.75 9.75 0 0 0 21.75 12a9.72 9.72 0 0 0-.832-3.948A9.752 9.752 0 0 0 12 2.25a9.735 9.735 0 0 0-7.8 3.9a.746.746 0 0 0-.09.12Zm.616 1.833a8.252 8.252 0 0 0 5.87 12.028l2.226-3.859a4.35 4.35 0 0 1-4.68-2.26L4.727 8.104Zm.934-1.382l2.228 3.855A4.352 4.352 0 0 1 12 7.65h7.011A8.245 8.245 0 0 0 12 3.75a8.233 8.233 0 0 0-6.34 2.97ZM19.744 9.15h-4.458A4.333 4.333 0 0 1 16.35 12a4.33 4.33 0 0 1-.675 2.329l-3.413 5.917A8.25 8.25 0 0 0 19.744 9.15ZM12 9.15a2.85 2.85 0 1 0 0 5.7a2.85 2.85 0 0 0 0-5.7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}