package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postdhl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.238 13.584a7.44 7.44 0 0 0-14.603 2.17"/><circle cx="23.07" cy="15.596" r="4.778" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.238 13.584a19.791 19.791 0 0 0 .639-4.214l5.833 1.443s-.771 12.223-13.622 12.223c-11.306 0-11.072-12.458-11.072-12.458l2.41-.632a15.822 15.822 0 0 0 1.21 5.808m1.484 9.111l-2.49 2.665m-8.376 7.603h5.036m-5.931 1.962h5.036m11.145-12.23l-2.49 2.665m7.534-2.665l2.488 2.665m1.753-2.665l2.488 2.665m5.152 7.603H43.5m-5.93 1.962h5.035m-30.052 2.752l4.143-8.589h1.396c2.362 0 3.362 1.933 2.223 4.295a7.93 7.93 0 0 1-6.367 4.294Zm12.678-8.589l-4.143 8.59m2.072-4.295h5.69m2.072-4.295l-4.144 8.59m7.204-8.59l-4.144 8.59h4.295M4.5 39.243h5.036m27.174 0h5.036"/>`),
		g.Group(children),
	)
}