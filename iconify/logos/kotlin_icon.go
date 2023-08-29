package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KotlinIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosKotlinIcon0" x1="99.991%" x2=".01%" y1="-.011%" y2="100.01%"><stop offset=".344%" stop-color="#E44857"/><stop offset="46.89%" stop-color="#C711E1"/><stop offset="100%" stop-color="#7F52FF"/></linearGradient></defs><path fill="url(#logosKotlinIcon0)" d="M256 256H0V0h256L128 127.949z"/>`),
		g.Group(children),
	)
}