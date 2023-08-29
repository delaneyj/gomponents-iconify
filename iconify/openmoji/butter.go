package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Butter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M19 33a3.416 3.416 0 0 0-3.314 2.983l-1.372 13.034A2.65 2.65 0 0 0 17 52h38a2.65 2.65 0 0 0 2.686-2.983l-1.372-13.034A3.416 3.416 0 0 0 53 33Z"/><path fill="#fff" d="M12.251 52h47.938L58 56H13.616l-1.365-4"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 33a3.416 3.416 0 0 0-3.314 2.983l-1.372 13.034A2.65 2.65 0 0 0 17 52h38a2.65 2.65 0 0 0 2.686-2.983l-1.372-13.034A3.416 3.416 0 0 0 53 33Zm-7.62 19H36"/><path d="M36 56H15a2.44 2.44 0 0 1-1.82-.583l-.239-.238l-3.934-8.009l-4.127-.019M60.62 52H36m0 4h21a2.44 2.44 0 0 0 1.82-.583l.239-.238l3.934-8.009l4.127-.019"/></g>`),
		g.Group(children),
	)
}