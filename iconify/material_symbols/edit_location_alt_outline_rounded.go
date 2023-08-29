package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLocationAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2q.675 0 1.338.113t1.287.312L13 4.075q-.25-.05-.488-.063T12 4Q9.475 4 7.737 5.738T6 10.2q0 1.775 1.475 4.063T12 19.35q3.05-2.8 4.525-5.088T18 10.2q0-.3-.025-.6t-.075-.575l1.65-1.65q.225.65.338 1.35T20 10.2q0 2.35-1.7 5.038t-4.975 5.587q-.275.25-.625.375t-.7.125q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2Zm0 8.2Zm6.35-6.35L17.2 2.7l-5.9 5.9q-.15.15-.225.338T11 9.325v.925q0 .325.213.538t.537.212h.925q.2 0 .388-.075t.337-.225l5.9-5.9l-.95-.95ZM20 4.1l.7-.7q.275-.275.275-.7T20.7 2l-.7-.7q-.275-.275-.7-.275t-.7.275l-.7.7L20 4.1Z"/>`),
		g.Group(children),
	)
}