package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.809 8.44c-.096-2.528 1.048-4.448 3.545-4.448H21.67c1.557 0 3.656 1.173 3.506 4.361c-.075 1.575-.275 5.277-.467 8.763A6.499 6.499 0 0 1 23.5 30h-15a6.499 6.499 0 0 1-1.227-12.88c-.21-3.66-.419-7.477-.464-8.68Zm3.473 13.302c-1.4 0-2.418-.94-2.856-2.27v3.598c0 1.495 1.343 2.677 2.85 2.677h11.446c1.508 0 2.851-1.192 2.851-2.677v-3.537c-.438 1.33-1.456 2.21-2.856 2.21H10.282Z"/>`),
		g.Group(children),
	)
}