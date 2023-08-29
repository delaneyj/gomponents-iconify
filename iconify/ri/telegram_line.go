package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm-3.11-8.83l-2.498-.779c-.54-.165-.543-.537.121-.804l9.733-3.76c.564-.23.886.061.703.79l-1.658 7.82c-.116.557-.451.69-.916.433l-2.551-1.888l-1.189 1.148c-.122.118-.221.219-.409.244c-.187.026-.341-.03-.454-.34l-.87-2.871l-.012.008Z"/>`),
		g.Group(children),
	)
}