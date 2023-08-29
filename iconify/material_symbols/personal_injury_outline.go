package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalInjuryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10q-1.65 0-2.825-1.175T8 6q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6q0 1.65-1.175 2.825T12 10Zm0-2q.825 0 1.413-.588T14 6q0-.825-.588-1.413T12 4q-.825 0-1.413.588T10 6q0 .825.588 1.413T12 8ZM4 22v-6.775q0-.85.425-1.562T5.6 12.55q1.275-.65 2.888-1.1T12 11q1.9 0 3.513.45t2.887 1.1q.75.4 1.175 1.113T20 15.225V20q0 .825-.587 1.413T18 22H9.75q-1.15 0-1.95-.8T7 19.25q0-1.15.8-1.95t1.95-.8h2.825l1.55-3.3q-.5-.1-1.025-.15T12 13q-1.8 0-3.2.437t-2.275.913q-.25.125-.388.363T6 15.225V22H4Zm5.75-2h1.2l.7-1.5h-1.9q-.3 0-.525.225T9 19.25q0 .3.225.525T9.75 20Zm3.4 0H18v-4.775q0-.275-.138-.512t-.362-.363q-.3-.15-.65-.312t-.75-.313L13.15 20ZM12 6Zm0 10.65Z"/>`),
		g.Group(children),
	)
}