package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.994 3.124a.75.75 0 0 1 .65-.374h6.713a.75.75 0 0 1 .649.374l6.643 11.482a.75.75 0 0 1-.01.77l-3.41 5.518a.75.75 0 0 1-.638.356H5.41a.75.75 0 0 1-.638-.356l-3.41-5.518a.75.75 0 0 1-.01-.77L7.994 3.124Zm.651 1.87L2.874 14.97l2.518 4.077l2.603-4.441l.001-.003l3.119-5.361l-2.47-4.246Zm3.338 5.739l-2.036 3.499h4.07l-2.034-3.5Zm3.77 3.499H20.7L14.924 4.25H9.947l5.806 9.982Zm4.902 1.5H9.073L6.718 19.75h11.455l2.482-4.018Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}