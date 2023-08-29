package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.536 42.272H9.464A3.964 3.964 0 0 1 5.5 38.308V24.213a3.963 3.963 0 0 1 1.163-2.805l14.535-14.52a3.963 3.963 0 0 1 5.604 0l14.535 14.52a3.963 3.963 0 0 1 1.163 2.805v14.095a3.964 3.964 0 0 1-3.964 3.964Zm-25.107-7.928H34.57v-8.488L24 15.296l-10.571 10.56ZM29.587 9.67L5.5 33.794m7.929.55v7.928m21.142-7.928v7.928"/>`),
		g.Group(children),
	)
}