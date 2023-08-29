package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMaleThreeHundredSixM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16 7a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V9.414l-3.11 3.11a6 6 0 1 1-1.414-1.414L18.587 8H17a1 1 0 0 1-1-1Zm0 9a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z" clip-rule="evenodd"/><path d="M18.4 36.8a3 3 0 1 0-4.8-3.6L11.5 36H9a3 3 0 0 0 0 6h4a3 3 0 0 0 2.4-1.2l3-4Zm22.86.415c-.474-.526-1.116-.822-1.786-.822H36.27l-.57-.262c-.428-.476-.935-1.116-1.486-1.813c-1.328-1.68-2.92-3.69-4.319-4.546A5.173 5.173 0 0 0 27.159 29H21.5l.586 13h4.73a5.332 5.332 0 0 0 2.272-.516a5.847 5.847 0 0 0 1.911-1.46l.57-.655c.224.249.442.598.67.961c.502.802 1.045 1.67 1.78 1.67h5.589c.67 0 1.178-.295 1.652-.821s.74-1.239.74-1.982c0-.744-.266-1.457-.74-1.982ZM33.5 28a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/></g>`),
		g.Group(children),
	)
}