package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditLocationAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-4.025-3.425-6.012-6.362T4 10.2q0-3.75 2.413-5.975T12 2q.675 0 1.338.113t1.287.312L13 4.075q-.25-.05-.488-.063T12 4Q9.475 4 7.737 5.738T6 10.2q0 1.775 1.475 4.063T12 19.35q3.05-2.8 4.525-5.088T18 10.2q0-.3-.025-.6t-.075-.575l1.65-1.65q.225.65.338 1.35T20 10.2q0 2.5-1.988 5.438T12 22Zm0-11.8Zm6.35-6.35L17.2 2.7L11 8.9V11h2.1l6.2-6.2l-.95-.95ZM20 4.1l.7-.7q.275-.275.275-.7T20.7 2l-.7-.7q-.275-.275-.7-.275t-.7.275l-.7.7L20 4.1Z"/>`),
		g.Group(children),
	)
}