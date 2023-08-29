package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedPaperLantern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.5 2A1.5 1.5 0 0 0 12 3.5V5h-.127a9.034 9.034 0 0 0-2.581 2h13.416a9.034 9.034 0 0 0-2.58-2H20V3.5A1.5 1.5 0 0 0 18.5 2h-5Zm-6.277 9c.17-.749.433-1.462.777-2.127V9h16v-.127c.344.665.607 1.378.777 2.127H7.223ZM7 15v-2h18v2H7Zm4.873 12a9.035 9.035 0 0 1-2.581-2h13.416a9.035 9.035 0 0 1-2.58 2H20v1.5a1.5 1.5 0 0 1-1.5 1.5h-5a1.5 1.5 0 0 1-1.5-1.5V27h-.127ZM8 23.127A8.936 8.936 0 0 1 7.223 21h17.554A8.935 8.935 0 0 1 24 23.127V23H8v.127ZM7 19v-2h18v2H7Z"/>`),
		g.Group(children),
	)
}