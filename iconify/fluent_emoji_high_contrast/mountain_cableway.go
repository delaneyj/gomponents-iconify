package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainCableway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastMountainCableway0)"><path d="M20.9 14h-9.8c-.61 0-1.1.49-1.1 1.1v8.8c0 .61.49 1.1 1.1 1.1h9.8c.61 0 1.1-.49 1.1-1.1v-8.8c0-.6-.49-1.1-1.1-1.1Z"/><path d="M29.023 2.88a1 1 0 0 1-.874 1.113L17 5.333V9h3.4c.33 0 .59.27.6.59V10h6.98c2.223 0 4 1.81 4.02 3.992l.005.52l-.005.009V26.99A5.014 5.014 0 0 1 26.989 32H5.01A5.014 5.014 0 0 1 0 26.99V14c.005-2.21 1.8-4 4.01-4H11v-.41c0-.32.26-.59.59-.59H15V5.574L4.11 6.883a1 1 0 0 1-.24-1.986l5.504-.661a1.01 1.01 0 0 1 .057-.009l12.48-1.5l.051-.005l5.949-.715a1 1 0 0 1 1.112.874ZM4.01 12A2.01 2.01 0 0 0 2 14h4.896C7.508 14 8 14.5 8 15.1v8.8c0 .61-.492 1.1-1.104 1.1H2v1.99C2 28.65 3.35 30 5.011 30H26.99c1.659 0 3.01-1.35 3.01-3.01V25h-4.896C24.492 25 24 24.51 24 23.9v-8.8c0-.61.492-1.1 1.104-1.1H30a2.032 2.032 0 0 0-2.02-2H4.01Z"/></g><defs><clipPath id="fluentEmojiHighContrastMountainCableway0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}