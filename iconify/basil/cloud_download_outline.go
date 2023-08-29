package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownloadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.086 12.531a.75.75 0 0 0-1.055-.117L12.75 13.44V10a.75.75 0 1 0-1.5 0v3.44l-1.282-1.026a.75.75 0 0 0-.937 1.172l2.498 1.998a.747.747 0 0 0 .946-.003l2.494-1.995a.75.75 0 0 0 .117-1.055Z"/><path fill="currentColor" fill-rule="evenodd" d="M8.46 7.284a5.296 5.296 0 0 1 9.734 2.543a4.386 4.386 0 1 1 .17 8.77H7A5.75 5.75 0 1 1 8.46 7.284Zm4.45-.922a3.796 3.796 0 0 0-3.422 2.15a.75.75 0 0 1-.947.372A4.25 4.25 0 1 0 7 17.096h11.362a2.887 2.887 0 1 0-.747-5.675a.75.75 0 0 1-.938-.812a3.795 3.795 0 0 0-3.769-4.247Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}