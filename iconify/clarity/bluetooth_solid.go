package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m26.52 24.52l-5.65-5.83l-1.46-1.5v-12l4.38 4.51l-3.6 3.71l2.24 2.29l4.09-4.22a2.54 2.54 0 0 0 0-3.56l-5.95-6.14a2.54 2.54 0 0 0-4.37 1.77v10.31l-5.53-5.7a1.6 1.6 0 1 0-2.3 2.23L15.75 18l-7 7.19a1.6 1.6 0 0 0 0 2.26a1.63 1.63 0 0 0 1.12.45a1.58 1.58 0 0 0 1.15-.49l5.11-5.27v10.31a2.53 2.53 0 0 0 1.59 2.36a2.44 2.44 0 0 0 .95.19a2.56 2.56 0 0 0 1.83-.77l5.95-6.15a2.54 2.54 0 0 0 .07-3.56Zm-7.12 6.31v-9.06l4.39 4.53Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}