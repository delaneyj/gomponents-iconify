package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Papirus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.545 4.509c-5.735.125-13.113 2.227-16.072 7.551c-.205 1.59.222 2.86 1.442 3.52c.523.11 1.22.551 1.271 1.173c2.04 5.626 1.744 11.613 2.207 17.523a27.403 27.403 0 0 0 .423 6.873c.08.341.574 1.803 1.016 1.928q3.815.922 5.595-.167c4.282-2.812 2.143-9.404 1.105-13.496c2.889-1.442 6.646-.736 9.324-2.851c11.236-7.27 7.878-22.478-6.311-22.054Zm-.957 8.532c3.098-.134 7.164 2.533 4.964 5.771c-1.907 2.355-5.813 3.041-8.753 2.632c-.533-2.347-.662-5.535.326-7.735a7.665 7.665 0 0 1 3.463-.668Z"/>`),
		g.Group(children),
	)
}