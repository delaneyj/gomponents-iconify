package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutomationTwoP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m95.847 148.193l52.206-52.594A31.855 31.855 0 0 1 144 80c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 17.673-14.327 32-32 32a31.85 31.85 0 0 1-16.89-4.815l-52.068 51.699A31.85 31.85 0 0 1 112 176c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32c0-17.673 14.327-32 32-32a31.854 31.854 0 0 1 15.847 4.193zM176 96c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16zm-96 96c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16z"/>`),
		g.Group(children),
	)
}