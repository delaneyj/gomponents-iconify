package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalAdvice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 14c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10h-9a1 1 0 0 1-1-1v-9ZM7.778 9A2.783 2.783 0 0 0 5 11.778C5 29.023 18.977 43 36.222 43A2.783 2.783 0 0 0 39 40.222v-6.204a2.783 2.783 0 0 0-2.778-2.778a19.28 19.28 0 0 1-6.028-.961a2.768 2.768 0 0 0-2.839.667l-3.389 3.389a25.94 25.94 0 0 1-10.302-10.3l3.39-3.39a2.784 2.784 0 0 0 .691-2.82l-.002-.007l-.002-.007a19.198 19.198 0 0 1-.963-6.033A2.783 2.783 0 0 0 14 9H7.778ZM33 19v-4h-4v-2h4V9h2v4h4v2h-4v4h-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}