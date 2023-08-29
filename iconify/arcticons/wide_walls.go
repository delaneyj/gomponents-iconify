package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WideWalls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="32.143" cy="32.316" r="4.229" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="8.457" height="26.552" x="31.478" y="10.724" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.229" ry="4.229" transform="rotate(23.198 35.707 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.133 11.797h0a4.228 4.228 0 0 0-5.552 2.221l-6.066 14.154c-.142.33-.387.605-.7.782h0a1.643 1.643 0 0 1-2.262-.66l-1.222-2.31a6.756 6.756 0 0 1 .056-6.426l.788-1.426l-.003-.001c.599-.84.902-1.905.743-3.047a4.238 4.238 0 0 0-3.71-3.604A4.23 4.23 0 0 0 4.5 15.683c0 .396.058.777.16 1.14l-.002.004l.007.017c.103.359.252.697.44 1.01l6.324 14.535a6.082 6.082 0 0 0 5.355 3.652l.05.002c.37.05.737.063 1.101.04l.056.002v-.004a5.65 5.65 0 0 0 4.797-3.41l6.566-15.321a4.228 4.228 0 0 0-2.221-5.553Z"/>`),
		g.Group(children),
	)
}