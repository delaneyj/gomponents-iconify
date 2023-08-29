package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLocationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.85 9.575L12.425 8.15L8.8 11.75q-.15.15-.225.338t-.075.387V13q0 .2.15.35t.35.15h.525q.2 0 .388-.075t.337-.225l3.6-3.625Zm.725-.725l.7-.7q.125-.125.125-.262t-.125-.263l-.9-.9q-.125-.125-.263-.125t-.262.125l-.7.7l1.425 1.425ZM12 21.325q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2q3.175 0 5.588 2.225T20 10.2q0 1.125-.438 2.363t-1.275 2.575Q17.45 16.475 16.2 17.9t-2.875 2.925q-.275.25-.625.375t-.7.125Z"/>`),
		g.Group(children),
	)
}