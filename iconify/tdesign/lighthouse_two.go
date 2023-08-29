package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LighthouseTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1v1.002h6V1h2v4h1v2h-1v3c0 .94.074 1.963.195 3H19v2h-1.523c.151.928.325 1.834.501 2.669a71.324 71.324 0 0 0 .9 3.758c.144.526.298 1.049.447 1.573H13.22l-1-4h-.439l-1 4H4.674c.15-.524.304-1.047.447-1.573a71.247 71.247 0 0 0 .9-3.758c.176-.835.35-1.741.501-2.669H5v-2h1.804c.121-1.037.196-2.06.196-3V7H6V5h1V1h2Zm0 3.002V5h6v-.998H9ZM15 7H9v3c0 .955-.07 1.974-.183 3h6.365A27.781 27.781 0 0 1 15 10V7Zm.451 8H8.548a60.8 60.8 0 0 1-.57 3.081A73.346 73.346 0 0 1 7.3 21h1.92l1-4h3.561l1 4H16.7a73.47 73.47 0 0 1-.678-2.919a60.8 60.8 0 0 1-.57-3.081ZM11 8.998h2.004v2.004H11V8.998Z"/>`),
		g.Group(children),
	)
}