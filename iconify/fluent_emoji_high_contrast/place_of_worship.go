package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceOfWorship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M24.486 13.37L16.6 8.247a.25.25 0 0 0-.271-.001L8.32 13.373a1 1 0 1 1-1.078-1.684l8.552-5.475a1.25 1.25 0 0 1 1.355.004l8.427 5.475a1 1 0 1 1-1.09 1.677Z"/><path d="M18 13a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-6.593 2.804a1.125 1.125 0 0 1 1.57.26l.888 1.239l1.905-1.18a1.375 1.375 0 0 1 1.95.549l2.088 4.12a1.375 1.375 0 0 1-.477 1.773l-1.725 1.122h2.363a1.125 1.125 0 1 1 0 2.25h-5.314c-1.368 0-1.896-1.781-.75-2.527l2.66-1.73l-1.299-2.599l-.913.565c-.615.38-1.42.22-1.841-.368l-1.364-1.904a1.125 1.125 0 0 1 .26-1.57Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}