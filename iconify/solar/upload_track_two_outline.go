package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTrackTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 1 0 1.85 18.315a.75.75 0 1 1 .3 1.47a10.82 10.82 0 0 1-2.15.215C6.063 22.75 1.25 17.937 1.25 12S6.063 1.25 12 1.25S22.75 6.063 22.75 12c0 .735-.074 1.454-.215 2.15a.75.75 0 0 1-1.47-.3A9.25 9.25 0 0 0 12 2.75ZM13 7c.414 0 .738.345.873.736A2.25 2.25 0 0 0 16 9.25a.75.75 0 0 1 0 1.5a3.734 3.734 0 0 1-2.25-.75v5a2.75 2.75 0 1 1-1.5-2.45v-4.8A.75.75 0 0 1 13 7Zm-.75 8a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0Zm5.22-.53a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06l-1.22-1.22V22a.75.75 0 0 1-1.5 0v-5.19l-1.22 1.22a.75.75 0 1 1-1.06-1.06l2.5-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}