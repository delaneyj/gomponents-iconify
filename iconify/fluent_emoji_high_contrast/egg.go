package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.085 4.837A8.773 8.773 0 0 1 16.331 1a8.772 8.772 0 0 1 7.202 3.78l.001.002a22.9 22.9 0 0 1 3.953 15.141v.003l-.09.925v.003A11.144 11.144 0 0 1 16.29 30.98c-5.743 0-10.555-4.36-11.105-10.082v-.003l-.08-.796v-.003A22.923 22.923 0 0 1 8.954 5.029l.002-.004l.129-.188ZM16.331 3a6.773 6.773 0 0 0-5.593 2.962l-.002.003l-.128.188a20.923 20.923 0 0 0-3.511 13.751l.08.8a9.153 9.153 0 0 0 9.114 8.276a9.144 9.144 0 0 0 9.114-8.311v-.005l.09-.927v-.002A20.9 20.9 0 0 0 21.89 5.919A6.772 6.772 0 0 0 16.33 3Z"/>`),
		g.Group(children),
	)
}