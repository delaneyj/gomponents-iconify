package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntrusionPrevention(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="22" cy="23.887" r="2" fill="currentColor"/><path fill="currentColor" d="M29.777 23.479A8.64 8.64 0 0 0 22 18a8.64 8.64 0 0 0-7.777 5.479L14 24l.223.521A8.64 8.64 0 0 0 22 30a8.64 8.64 0 0 0 7.777-5.479L30 24zM22 28a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm3-18H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h21a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 4v4h21V4zm8 24H4v-4h8v-2H4a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h8z"/><path fill="currentColor" d="M28 12H7a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h5v-2H7v-4h21v2h2v-2a2.002 2.002 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}