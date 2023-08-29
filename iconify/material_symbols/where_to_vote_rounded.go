package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhereToVoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.95 10.5l-.7-.7q-.3-.3-.7-.3t-.7.3q-.3.3-.3.713t.3.712l1.4 1.425q.3.3.7.3t.7-.3L15.2 9.1q.3-.3.3-.712t-.3-.713q-.3-.3-.712-.3t-.713.3L10.95 10.5Zm9.05-.3q0 1.125-.438 2.363t-1.275 2.574Q17.45 16.476 16.2 17.9t-2.875 2.925q-.275.25-.625.375t-.7.125q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2q3.175 0 5.588 2.225T20 10.2Z"/>`),
		g.Group(children),
	)
}