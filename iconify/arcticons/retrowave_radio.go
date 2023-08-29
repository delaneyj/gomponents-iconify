package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RetrowaveRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.41 26.471a15.044 15.044 0 0 0-1.114 4.663c0 2.003.901 4.055.901 4.606s-3.104 3.955-3.905 4.456a14.339 14.339 0 0 1-3.304.9a5.03 5.03 0 0 0-2.252 2.404c-.35-1.565-.463-2.053-1.29-2.69s-1.507-.876-1.432-1.326M32.81 9.053a28.839 28.839 0 0 0-9.474-1.85c-3.454 0-7.56.801-7.56 5.758c-2.353 1.451-2.303 6.358-2.303 6.358"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.185 7.395A5.368 5.368 0 0 1 30.72 4.5c2.528 0 5.93 1.124 6.108 7.434c.065 2.34-.075 3.292 0 4.218s.818 3.567.096 5.72s-3.776 5.582-6.93 4.681s-3.73-4.656-3.204-6.959s2.403-6.333 6.158-5.432c0-1.552.75-7.159-2.103-7.159a1.576 1.576 0 0 0-1.59.974"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.1 17.742c.576.776 2.578 3.705.576 5.632c-1.963 1.89-3.255 1.427-3.255 1.427s-1.701-2.503-.65-4.381a6.834 6.834 0 0 1 3.329-2.678Zm-6.913 6.295c-.68 2.377 1.469 6.146 1.469 8.474c0 1.314-1.124 2.71-3.004 3.38c-1.69.6-4.018 3.867-7.284 3.867s-6.459-3.154-6.459-6.196s1.39-6.082 3.492-6.082s2.403 1.802 3.868 1.802s4.205-.939 4.543-3.567a16.614 16.614 0 0 1 1.164-4.419"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.925 21.822c-.85.951-2.503 3.75-2.503 4.879s1.511 1.329 1.511 1.329m6.199-6.208c0 2.578-.976 3.73-2.178 3.73s-2.803-1.602-2.803-3.905c0-1.08.15-1.527.976-1.527h2.754c.425 0 1.251.958 1.251 1.702ZM19.5 20.5h7.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.41 22.483c.04-.686.015-2.063-.21-2.213s-2.203-1.226-2.954-.951s-.55 3.23.075 4.105a5.708 5.708 0 0 0 1.622 1.589m2.376-4.513H14.31"/>`),
		g.Group(children),
	)
}