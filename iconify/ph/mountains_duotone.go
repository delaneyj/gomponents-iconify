package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M94.9 75.94L125.65 128h-75.3L81.1 75.94a8 8 0 0 1 13.8 0ZM164 72a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm29.44 35.92a8 8 0 0 0-13.77 0l-33.06 55.79L168 200h80Z" opacity=".2"/><path d="M164 80a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm0-40a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm90.88 155.92l-54.56-92.08A15.87 15.87 0 0 0 186.55 96a15.85 15.85 0 0 0-13.76 7.84L146.63 148l-44.84-76.1a16 16 0 0 0-27.58 0L1.11 195.94A8 8 0 0 0 8 208h240a8 8 0 0 0 6.88-12.08ZM88 80l23.57 40H64.43ZM22 192l33-56h66l18.74 31.8L154 192Zm150.57 0l-16.66-28.28L186.55 112L234 192Z"/></g>`),
		g.Group(children),
	)
}