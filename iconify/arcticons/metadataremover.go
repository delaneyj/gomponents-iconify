package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metadataremover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.15 21.93l2.93-2.72l3.26 3.21L27 13.51l9.85 9.43V7.21h-25.7Zm8.55-12a2.92 2.92 0 1 1-2.87 3v-.13a2.9 2.9 0 0 1 2.87-2.9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Zm-6.72 37.13a6.09 6.09 0 0 1-6.17-6.08a6.17 6.17 0 1 1 6.17 6.08ZM40 23.88a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2m6.83 7.13h6.43m0 4.13h-6.43m0 3.86h6.51m7.11-6.91l5.73 5.73m-5.75.04l5.73-5.73"/><path fill="currentColor" d="M11.9 39.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Zm0-3.86a.75.75 0 1 1 0-1.5a.75.75 0 0 1 .75.74a.77.77 0 0 1-.75.76Zm0-4.14a.75.75 0 0 1-.75-.75a.74.74 0 0 1 .75-.74a.75.75 0 0 1 .75.75a.77.77 0 0 1-.75.74Z"/>`),
		g.Group(children),
	)
}