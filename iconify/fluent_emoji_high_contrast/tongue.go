package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 17.723c3.057-2.569 5-6.42 5-10.723H2c0 4.31 1.95 8.168 5.016 10.736v5.201c0 1.44.595 3.436 2.006 5.084C10.464 29.707 12.737 31 16 31c3.262 0 5.539-1.293 6.985-2.977C24.4 26.375 25 24.378 25 22.938v-5.215ZM9.016 14.969c0-.509.393-.993 1.374-1.396c.927-.38 1.995-.51 2.422-.51h6.5c.645 0 1.668.165 2.496.55c.852.397 1.192.864 1.192 1.356v7.969c0 .965-.431 2.5-1.532 3.782C20.398 27.965 18.675 29 16 29s-4.394-1.035-5.46-2.28c-1.096-1.28-1.524-2.815-1.524-3.782v-7.97Z"/>`),
		g.Group(children),
	)
}