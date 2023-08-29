package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeMinimalisticOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.652 2.469a.75.75 0 0 1 1.061.002l6.819 6.849a.75.75 0 0 1-1.064 1.058l-.23-.231L10.02 20.411a4.54 4.54 0 0 1-6.438 0a4.58 4.58 0 0 1 0-6.457l2.682-2.695a.5.5 0 0 1 .01-.01l7.531-7.564l-.155-.156a.75.75 0 0 1 .002-1.06Zm1.212 2.28L8.418 11.22a2.956 2.956 0 0 1 2.345 2.575c.058.522.39.971.867 1.178l1.192.495l6.357-6.385l-4.316-4.336Zm-3.185 11.87l-.637-.265a2.96 2.96 0 0 1-1.769-2.394a1.455 1.455 0 0 0-1.281-1.289l-.915-.102l-2.432 2.443a3.08 3.08 0 0 0 0 4.34a3.039 3.039 0 0 0 4.312 0l2.722-2.734Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}