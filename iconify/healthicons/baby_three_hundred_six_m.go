package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyThreeHundredSixM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.8 26.6a3 3 0 0 1 .6 4.2l-3 4A3 3 0 0 1 13 36H9a3 3 0 1 1 0-6h2.5l2.1-2.8a3 3 0 0 1 4.2-.6Z" clip-rule="evenodd"/><path d="M39.474 30.393H36.27l-.57-.262c-1.46-1.622-3.824-5.148-5.805-6.359A5.172 5.172 0 0 0 27.159 23H21.5l.586 13h4.729a5.332 5.332 0 0 0 2.273-.516a5.847 5.847 0 0 0 1.911-1.46l.57-.655c.719.799 1.383 2.631 2.45 2.631h5.589c.67 0 1.178-.295 1.652-.821s.74-1.239.74-1.982c0-.744-.266-1.457-.74-1.982c-.474-.526-1.116-.822-1.786-.822ZM40 15.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Z"/><path fill-rule="evenodd" d="M33.5 20a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0 2a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}