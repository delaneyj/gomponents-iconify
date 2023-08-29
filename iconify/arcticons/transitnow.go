package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transitnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="17.07" cy="35.78" r="2.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.64 9.98h22.72v4.63H12.64z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.92 4.5H16.08a7.58 7.58 0 0 0-7.58 7.58v19a3 3 0 0 0 .36 1.43l4.78 9a3.78 3.78 0 0 0 3.34 2H31a3.78 3.78 0 0 0 3.34-2l4.78-9a3 3 0 0 0 .36-1.43v-19a7.58 7.58 0 0 0-7.56-7.58Z"/><circle cx="30.93" cy="35.78" r="2.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.36 14.61v9.76l-3.61 4.66h-15.5l-3.61-4.66v-9.76h22.72z"/>`),
		g.Group(children),
	)
}