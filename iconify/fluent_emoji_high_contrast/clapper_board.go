package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.16 2l-4.21 1.12l-14.7 3.89l-1.48.4l-1.28.34L4 8.14l.99 3.73l2.76-.73l8.706-2.313l.004.003l5.7-1.51l-.006-.005L26.66 6.12l1.49-.39L27.16 2ZM9.211 7.79l4.505-1.192l1.781 1.449L11 9.24L9.211 7.79Zm10.201-2.7l3.512-.928l.282-.075h.001l.706-.188l1.784 1.442l-4.504 1.195l-1.78-1.445ZM29 28V13H5v15a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2Zm-2.475-14l-1.995 2.74h-4.65L21.875 14h4.65Zm-15.2 0h4.66l-1.995 2.74H9.33L11.325 14Z"/>`),
		g.Group(children),
	)
}