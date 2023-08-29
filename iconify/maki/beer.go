package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 5V2s-1-1-4.5-1S3 2 3 2v3a9.27 9.27 0 0 0 1 4a5.63 5.63 0 0 1 0 4.5s0 1 3.5 1s3.5-1 3.5-1A5.63 5.63 0 0 1 11 9a9.27 9.27 0 0 0 1-4zm-4.5 8.5a7.368 7.368 0 0 1-2.36-.28c.203-.722.304-1.47.3-2.22h4.12c-.004.75.097 1.498.3 2.22a7.368 7.368 0 0 1-2.36.28zm0-8.5A10.65 10.65 0 0 1 4 4.5v-2A10.74 10.74 0 0 1 7.5 2a10.74 10.74 0 0 1 3.5.5v2c-1.13.36-2.314.53-3.5.5z"/>`),
		g.Group(children),
	)
}