package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.325q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-.8.125-1.525t.35-1.375l-3.1-3.1q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.713t-.3.712q-.3.3-.712.3t-.713-.3l-4.1-4.1q-.625.65-1.25 1.288t-1.1 1.037q-.275.25-.625.375t-.7.125Zm6.2-6L13.775 10.9q.125-.2.175-.425T14 10q0-.825-.587-1.413T12 8q-.25 0-.475.05t-.425.175L6.775 3.9Q7.85 2.975 9.2 2.488T12 2q3.175 0 5.588 2.225T20 10.2q0 1.2-.45 2.463t-1.35 2.662Z"/>`),
		g.Group(children),
	)
}