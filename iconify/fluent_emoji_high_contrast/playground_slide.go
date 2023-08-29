package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaygroundSlide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastPlaygroundSlide0)"><path d="M3.5 14.047h3.314l-1.148-2.64A4 4 0 0 0 3.5 9.294v4.754Zm0 4.969h5.483a1.591 1.591 0 0 1-.052-.1l-1.24-2.852H3.5v2.952Zm0 4.968h5.984V21.03H3.5v2.953Z"/><path d="M10.836 1a4.652 4.652 0 0 0-3.082 1.164A4.669 4.669 0 0 0 0 5.672V29a2 2 0 0 0 2 2h1.5a2 2 0 0 0 2-2v-1h1.984v1a2 2 0 0 0 2 2h1.5a2 2 0 0 0 2-2v-5.872l3.155 6.722A2 2 0 0 0 17.95 31H30a2 2 0 0 0 1.816-2.837L27.267 18.3a6.519 6.519 0 0 0-5.9-3.78h-.136l-1.953-4.5a7.51 7.51 0 0 0-3.767-3.837v-.5A4.676 4.676 0 0 0 10.836 1Zm0 2a2.672 2.672 0 0 1 2.672 2.672v1.942a5.5 5.5 0 0 1 3.932 3.2l2.476 5.7h1.448a4.5 4.5 0 0 1 4.087 2.616L30 29h-1.652L24.2 20l-.111-.24a3 3 0 0 0-2.725-1.744h-1.776A.991.991 0 0 1 19.41 18a1 1 0 0 1-.74-.586l-2.607-6.006A4 4 0 0 0 12.394 9h-.363v-.016l-.023.012V5.672a1.172 1.172 0 0 0-2.344 0v3.312l-1.5-.001V5.672A2.672 2.672 0 0 1 10.836 3ZM7.344 8.982l-1.5-.001V5.67a1.172 1.172 0 0 0-2.344 0v2.04c.85.242 1.629.68 2.271 1.287a5.509 5.509 0 0 1 1.27 1.815L10.17 18l.007.016v.002h1.448a4.928 4.928 0 0 1 1.99.593c.994.586 1.843 1.533 2.294 2.511L19.601 29H17.95l-3.4-7.25a3.82 3.82 0 0 0-2.922-2.234h-.641V29h-1.5v-3H3.5v3H2V5.672a2.672 2.672 0 0 1 5.344 0v3.31Z"/></g><defs><clipPath id="fluentEmojiHighContrastPlaygroundSlide0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}