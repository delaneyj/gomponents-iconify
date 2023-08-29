package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOnFire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.036 2.684C16.29 5.254 17.89 6.59 19 6.97c.5.17.99-.474 1-.996V5c0-.542.63-1.16 1.139-.96C25.814 5.857 30 9.75 30 15c0 7.64-11.48 14.775-13.649 15.91a.729.729 0 0 1-.702 0C13.489 29.775 2 22.64 2 15c0-4.457 1.74-8.434 5-9.9c.755-.339.95.349 1.139 1.02c.097.344.192.683.361.88c.336.392.829.111 1.063-.36c.784-1.567 2.724-4.719 5.698-4.638c.398.01.734.3.775.682Zm-.038 9.776s1.71-2.68 3.938-3.08c5.51-.994 7.788 3.943 6.864 7.617C25.16 23.53 15.998 29.25 15.998 29.25S6.836 23.53 5.196 16.997c-.913-3.674 1.356-8.611 6.865-7.618c2.229.402 3.937 3.081 3.937 3.081Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}