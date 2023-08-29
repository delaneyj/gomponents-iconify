package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.57 1l6.2 4l.23.38v9.2l-.76.42L8 11l-6.24 4l-.76-.42v-9.2L1.23 5l6.2-4h1.14zm-.06 9.13L14 13.67V5.65l-5.49-3.5V5h-1V2.13L2 5.67v8l5.51-3.56v.02h1zm.9-4.78l.71-.7l2.47 2.48v.71l-2.46 2.46l-.7-.7L11.02 8h-6L6.6 9.6l-.7.7l-2.46-2.46v-.71l2.48-2.48l.7.7L4.98 7h6.08L9.41 5.35z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}