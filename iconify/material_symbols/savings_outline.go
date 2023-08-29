package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SavingsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11q.425 0 .713-.288T17 10q0-.425-.288-.713T16 9q-.425 0-.713.288T15 10q0 .425.288.713T16 11ZM8 9h5V7H8v2ZM4.5 21q-.85-2.85-1.675-5.688T2 9.5q0-2.3 1.6-3.9T7.5 4h5q.725-.95 1.763-1.475T16.5 2q.625 0 1.063.438T18 3.5q0 .15-.038.3t-.087.275q-.1.275-.188.563t-.137.587L19.825 7.5H22v6.975l-2.825.925L17.5 21H12v-2h-2v2H4.5ZM6 19h2v-2h6v2h2l1.55-5.15l2.45-.825V9.5h-1L15.5 6q0-.5.063-.963t.187-.937q-.725.2-1.275.688T13.675 6H7.5Q6.05 6 5.025 7.025T4 9.5q0 2.45.675 4.788T6 19Zm6-7.45Z"/>`),
		g.Group(children),
	)
}