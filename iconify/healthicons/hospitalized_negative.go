package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalizedNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsHospitalizedNegative0)"><path d="m18.117 21.188l.364-.364a1.77 1.77 0 0 0 0-2.504l-1.8-1.8a1.771 1.771 0 0 0-2.504-.002l-.35.35l4.29 4.32ZM36 31v3H11v-8.644l5.275 5.311c.212.213.498.333.797.333H36Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm29 8V6h2v2a1 1 0 0 0 1 1h6a2 2 0 0 1 2 2v12.578a3.97 3.97 0 0 1 2 3.455C42 29.223 40.239 31 38.066 31H38v3h4v2h-3.05a3.5 3.5 0 1 1-4.899 0H13.948a3.5 3.5 0 1 1-4.899 0H6v-2h3V23.343l-.833-.839a4.034 4.034 0 0 1 0-5.676a3.953 3.953 0 0 1 3.892-1.021l.704-.704a3.771 3.771 0 0 1 5.333.002l1.8 1.801a3.77 3.77 0 0 1-.002 5.333l-.368.368l.366.369c.058.058.137.09.218.09H38V11h-5v1h1a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h1v-1.17A3.001 3.001 0 0 1 29 8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHospitalizedNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}