package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.525q-.325 0-.638-.112t-.587-.313l-6.75-5.25q-.4-.3-.388-.787t.413-.788q.275-.2.6-.2t.6.2L12 18.5l6.75-5.225q.275-.2.6-.2t.6.2q.4.3.413.787t-.388.788l-6.75 5.25q-.275.2-.587.313t-.638.112Zm0-5.05q-.325 0-.638-.112t-.587-.313L4.025 9.8q-.2-.15-.3-.363T3.625 9q0-.225.1-.438t.3-.362l6.75-5.25q.275-.2.588-.313T12 2.525q.325 0 .638.113t.587.312l6.75 5.25q.2.15.3.363t.1.437q0 .225-.1.438t-.3.362l-6.75 5.25q-.275.2-.587.313t-.638.112Z"/>`),
		g.Group(children),
	)
}