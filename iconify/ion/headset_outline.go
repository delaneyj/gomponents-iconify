package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M83 384c-13-33-35-93.37-35-128C48 141.12 149.33 48 256 48s208 93.12 208 208c0 34.63-23 97-35 128"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="m108.39 270.13l-13.69 8c-30.23 17.7-31.7 72.41-3.38 122.2s75.87 75.81 106.1 58.12l13.69-8a16.16 16.16 0 0 0 5.78-21.87L130 276a15.74 15.74 0 0 0-21.61-5.87Zm295.22 0l13.69 8c30.23 17.69 31.74 72.4 3.38 122.19s-75.87 75.81-106.1 58.12l-13.69-8a16.16 16.16 0 0 1-5.78-21.87L382 276a15.74 15.74 0 0 1 21.61-5.87Z"/>`),
		g.Group(children),
	)
}