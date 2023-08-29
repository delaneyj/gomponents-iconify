package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.85 0-3.488-.713T5.65 19.35q-1.225-1.225-1.938-2.863T3 13q0-1.5.463-2.888T4.8 7.6L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-1.7-1.7q-1.2.875-2.587 1.338T12 22ZM10 1h4q.425 0 .713.288T15 2q0 .425-.288.713T14 3h-4q-.425 0-.713-.288T9 2q0-.425.288-.713T10 1Zm2 3q1.5 0 2.938.5t2.712 1.45l.7-.7q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-.7.7q.95 1.275 1.45 2.713T21 13q0 1.05-.263 2.087t-.787 2.063L13 10.2V9q0-.425-.288-.713T12 8q-.25 0-.463.1t-.337.3L7.825 5.025q.95-.5 2.013-.763T12 4Z"/>`),
		g.Group(children),
	)
}