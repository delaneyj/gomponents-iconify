package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaleCondom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m27.886 6l4.108 1.14l-.732 2.85l2.733.764l-.732 2.85l2.733.763l-.732 2.85l2.734.764l-.733 2.85l2.734.763l-.733 2.85l2.734.764l-1.107 4.28L20.105 42l-4.1-1.145l.733-2.85l-2.734-.764l.732-2.85l-2.733-.763l.732-2.85l-2.733-.764l.732-2.85l-2.733-.763l.732-2.85L6 22.787l1.09-4.27L27.886 6ZM16.207 28.5a9 9 0 1 0 15.589-9a9 9 0 0 0-15.589 9Zm1.735-1.002A6.999 6.999 0 1 0 30.064 20.5a6.999 6.999 0 0 0-12.122 6.998Zm8.558.831a5 5 0 1 1-5-8.658a5 5 0 0 1 5 8.658Zm-.909-1.575a3.18 3.18 0 1 1-3.18-5.51a3.18 3.18 0 0 1 3.18 5.51Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}