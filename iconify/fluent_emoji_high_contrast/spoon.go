package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26.243 29.461l-12.88-13.615a3.223 3.223 0 0 0-2.618-.98c-2.114.172-4.562-.816-6.446-2.808C1.419 9.014.962 4.775 3.27 2.592C5.58.408 9.785 1.1 12.665 4.145c1.884 1.992 2.734 4.49 2.452 6.598c-.136.978.154 1.954.826 2.663L28.85 27.053a1.767 1.767 0 0 1-.126 2.537c-.728.642-1.829.562-2.482-.129ZM5.337 11.076c2.162 2.287 5.324 2.808 7.06 1.165c1.738-1.643 1.393-4.828-.77-7.114C9.464 2.84 6.303 2.318 4.566 3.96c-1.737 1.643-1.392 4.829.77 7.115Z"/>`),
		g.Group(children),
	)
}