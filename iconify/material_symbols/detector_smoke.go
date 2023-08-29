package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorSmoke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.1 8l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4ZM13 22.3l-1.9-.6l.4-1.3q.15-.6.1-1.163t-.4-1.137q-.625-.95-.787-1.963T10.6 14l.4-1.3l1.9.6l-.4 1.3q-.2.575-.125 1.188T12.8 16.9q.575.95.763 2T13.4 21l-.4 1.3Zm-4.4 0l-1.9-.6l.4-1.3q.15-.575.1-1.175T6.8 18.1q-.65-.925-.8-1.975T6.2 14l.4-1.3l1.9.6l-.4 1.3q-.2.575-.125 1.2t.425 1.1q.6.9.775 1.988T9 21l-.4 1.3Zm8.7 0l-1.9-.6l.4-1.3q.15-.575.1-1.175t-.4-1.125q-.65-.925-.8-1.975T14.9 14l.4-1.3l1.9.6l-.4 1.3q-.2.6-.125 1.187T17.1 16.9q.6.95.763 2T17.7 21l-.4 1.3Z"/>`),
		g.Group(children),
	)
}