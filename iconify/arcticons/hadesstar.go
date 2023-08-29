package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hadesstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="13.3" fill="none" stroke="currentColor" stroke-dasharray="1 2" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.4 38.63a16 16 0 0 0 4.56.69a15.34 15.34 0 0 0 10.86-26.19a17 17 0 0 0-3.54-2.64l1.46-2.7A19.62 19.62 0 0 1 37 11a18.35 18.35 0 0 1 5.39 13h3.07A21.48 21.48 0 0 0 24 2.47v3.07a20.07 20.07 0 0 1 5.47.82l-.95 2.94A16.39 16.39 0 0 0 24 8.61A15.34 15.34 0 0 0 13.1 34.8a16.38 16.38 0 0 0 3.54 2.64l-1.46 2.71A19.74 19.74 0 0 1 10.93 37a18.35 18.35 0 0 1-5.39-13H2.47A21.49 21.49 0 0 0 24 45.47V42.4a20.14 20.14 0 0 1-5.48-.83Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.07 31.06l.33 5.43l-4.13 4l-5.2-.49m10.22-13.77a17.89 17.89 0 0 1-1.23 4.85a17.06 17.06 0 0 1-9 9a17.86 17.86 0 0 1-4.85 1.24"/><circle cx="24" cy="24" r="11.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}