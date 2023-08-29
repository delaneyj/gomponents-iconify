package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.679 2.679 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68H9.5M30.434 17v5.524h3.609m0 7.085h-3.609v1.153a8.239 8.239 0 1 1-8.238-8.238h1.153V17m-1.153 12.61a1.153 1.153 0 1 0 1.153 1.152V29.61Zm11.847-3.544h-1.288m-5.864-2.537v-2.547m0 7.802v-2.221"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.72 33.864a4.682 4.682 0 0 0 1.073-2.142m-5.935 3.542a4.73 4.73 0 0 0 2.49.05m-5.783-3.768a4.677 4.677 0 0 0 1.022 2.22m.983-6.898a4.716 4.716 0 0 0-1.614 1.872m5.393-2.674h-1.153m4.695-7.778V17m-3.542 5.524v7.085m7.085-7.085v7.085"/>`),
		g.Group(children),
	)
}