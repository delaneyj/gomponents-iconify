package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayForWorkRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-2.25 0-3.925-1.413t-2-3.587q-.05-.4.225-.7T7 14q.425 0 .713.213t.362.562Q8.35 16.2 9.463 17.1T12 18q1.425 0 2.538-.9t1.387-2.325q.075-.35.363-.562T17 14q.425 0 .7.3t.225.7q-.325 2.175-2 3.588T12 20Zm-1-8.85V6q0-.425.288-.713T12 5q.425 0 .713.288T13 6v5.15l.9-.9q.3-.3.7-.288t.7.313q.275.3.287.7t-.287.7l-2.6 2.6q-.15.15-.325.213T12 14.55q-.2 0-.375-.063t-.325-.212l-2.6-2.6q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3l.875.875Z"/>`),
		g.Group(children),
	)
}