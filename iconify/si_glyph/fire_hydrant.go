package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHydrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 14.047h-.326c-.371 0-.674.457-.674 1.019v.893h7.987v-.893c0-.563-.303-1.019-.674-1.019h-.357V6H5v8.047zM4 4v.942L12 5V4H4zm3-2.858c-1.148.27-1.969.993-1.969 1.842h5.938c0-.854-.83-1.58-1.988-1.847V0H7v1.142zM3 8V7h.947v1h.027v.97h-.027v.938H3V8.97H2V8h1zm9.953.938v.968H12V7h.953v1h.989v.938h-.989z"/>`),
		g.Group(children),
	)
}