package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.023 30.642h24V43.5h-24zm12-20.564c-3.707 3.404-4.653 8.746-1.627 20.564"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.21 30.642c-2.321-9.32 4.67-17.747 10.644-13.426c5.374 3.888-1.85 13.426-10.643 13.426Zm-6.42 0c2.321-9.32-4.67-17.747-10.644-13.426c-5.374 3.888 1.85 13.426 10.643 13.426Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.023 10.078c3.708 3.404 4.653 8.746 1.628 20.564m-1.628-20.564V4.5m-1.686 1.796h3.372M14.758 37.071h18.53m-18.53 3.214h18.53m-18.53-6.429h18.53"/>`),
		g.Group(children),
	)
}