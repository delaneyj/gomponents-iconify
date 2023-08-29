package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScVimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M38 19.6c-.1 2.7-2 6.4-5.6 11.1c-3.8 4.9-7 7.4-9.6 7.4c-1.6 0-3-1.5-4.1-4.5c-.7-2.7-1.5-5.5-2.2-8.2c-.8-3-1.7-4.5-2.7-4.5c-.2 0-.9.4-2.2 1.3l-1.3-1.7c1.4-1.2 2.7-2.4 4-3.6c1.8-1.6 3.2-2.4 4.1-2.5c2.2-.2 3.5 1.3 4 4.4c.5 3.4.9 5.5 1.1 6.4c.6 2.8 1.3 4.2 2.1 4.2c.6 0 1.5-.9 2.6-2.8c1.2-1.8 1.8-3.2 1.9-4.2c.2-1.6-.5-2.4-1.9-2.4c-.7 0-1.3.2-2 .5c1.4-4.5 4-6.6 7.8-6.5c2.8.1 4.2 1.9 4 5.6z"/>`),
		g.Group(children),
	)
}