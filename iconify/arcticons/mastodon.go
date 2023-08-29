package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mastodon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.066 20.401c0 6.258-.972 8.398-2.884 10.19C34.308 34.226 24 33.943 24 33.943a40.292 40.292 0 0 1-7.014-.528s-1.126 7.414 13.104 3.778l-.144 4.24c-2.169.29-22.808 7.92-22.957-18.656L6.933 20.4c0-6.258.05-8.482 2.884-11.611C13.383 4.853 24 5.355 24 5.355s10.617-.502 14.182 3.435c2.834 3.129 2.884 5.354 2.884 11.611"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 23.313v-5.978a4.654 4.654 0 0 0-4.654-4.654h0a4.654 4.654 0 0 0-4.654 4.654v9.646M24 17.335a4.654 4.654 0 0 1 4.654-4.654h0a4.654 4.654 0 0 1 4.654 4.654v9.646"/>`),
		g.Group(children),
	)
}