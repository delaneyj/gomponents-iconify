package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourKBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Zm4.25-4a.75.75 0 0 0-1.5 0v2a2.75 2.75 0 0 0 2.75 2.75h2.25V16a.75.75 0 0 0 1.5 0V8a.75.75 0 0 0-1.5 0v3.25H7.5c-.69 0-1.25-.56-1.25-1.25V8Zm12.77-.54a.75.75 0 0 1 .02 1.06l-2.666 2.773l2.757 4.302a.75.75 0 1 1-1.262.81l-2.564-4l-1.055 1.097V16a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 1.5 0v3.338l3.71-3.858a.75.75 0 0 1 1.06-.02Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}