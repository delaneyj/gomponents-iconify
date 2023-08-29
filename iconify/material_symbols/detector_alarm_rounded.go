package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorAlarmRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.425 0-.713-.288T11 20v-3q0-.425.288-.713T12 16q.425 0 .713.288T13 17v3q0 .425-.288.713T12 21Zm7.775-3.225q-.275.275-.7.275t-.7-.275l-2.15-2.125q-.3-.3-.3-.7t.3-.7q.3-.3.713-.3t.712.3l2.125 2.125q.275.275.288.688t-.288.712Zm-15.55 0q-.275-.275-.275-.7t.275-.7l2.125-2.15q.3-.3.7-.3t.7.3q.3.3.3.713t-.3.712l-2.125 2.125q-.275.275-.688.288t-.712-.288ZM8.1 8l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4Z"/>`),
		g.Group(children),
	)
}