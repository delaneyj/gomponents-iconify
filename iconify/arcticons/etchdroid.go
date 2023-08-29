package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etchdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.25 13.716l2.59-2.551a3.022 3.022 0 0 0-4.282-4.266l-.002.002l-2.55 2.599a18.41 18.41 0 0 0-21.515 0l-2.57-2.57a3.174 3.174 0 0 0-2.171-.915a2.979 2.979 0 0 0-2.122.876a3.018 3.018 0 0 0 0 4.264l2.599 2.55a18.498 18.498 0 0 0-3.476 10.807v2.738h36.994v-2.738a18.497 18.497 0 0 0-3.495-10.796Zm-19.87 6.98a2.92 2.92 0 1 1-2.92-2.92a2.92 2.92 0 0 1 2.92 2.92Zm12.656 2.92a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.92 2.92ZM12.341 27.25H38.15v14.737H12.341zm12.904 0v14.737"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.938 32.888h5.467v3.461h-5.467zm-12.826 0h5.467v3.461h-5.467z"/>`),
		g.Group(children),
	)
}