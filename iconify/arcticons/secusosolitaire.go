package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Secusosolitaire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.88 16.6c-3.19-3-4.62-3.43-4.62-6.07a2.35 2.35 0 0 1 2.34-2.36h.05a2.59 2.59 0 0 1 2.23 1.36a2.63 2.63 0 0 1 2.24-1.36a2.35 2.35 0 0 1 2.39 2.31v.08c0 2.61-1.44 3-4.63 6.04ZM32.26 40l-2.5-4.23l2.53-4.23l2.47 4.23Zm-15.37-9.51a1.84 1.84 0 1 1-1.84 1.84a1.84 1.84 0 0 1 1.84-1.84ZM14 37.58a1.84 1.84 0 1 1 1.83-1.85v0A1.83 1.83 0 0 1 14 37.58Zm1.46 2.12l1.41-2.48l1.44 2.48Zm4.17-2.12a1.84 1.84 0 1 1 1.84-1.84h0a1.83 1.83 0 0 1-1.83 1.83Zm11.63-20.95l1-1.8l1 1.8ZM34.05 15a2.11 2.11 0 0 1-1.8-1.1a2.12 2.12 0 0 1-1.8 1.1a1.9 1.9 0 0 1-1.94-1.86v-.07c0-2.13 1.16-2.45 3.74-4.9c2.57 2.45 3.73 2.77 3.73 4.9A1.9 1.9 0 0 1 34.13 15Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 43.5h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}