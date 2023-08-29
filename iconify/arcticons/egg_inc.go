package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggInc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.756 43.16l-2.29-5.01l-.095-7.025m.527-.19l1.855-.611l1.159-2.683l-1.103-2.665l-3.641 2.247l-.668 2.215l.544 2.108Zm1.911-5.959a6.553 6.553 0 0 1 3.146-.153l-3.077-4.495Zm.069-4.648l-3.634-6.399l-5.359-2.28l-4.633-1.167l2.35 5.752l7.27.367Zm-11.276-4.094l-7.515 4.484l-5.486 9.325l-4.66 3.908"/><circle cx="24.4" cy="21.149" r="1.918" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="21.247" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}