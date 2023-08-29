package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignjustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.865 512h-896q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-192h-896q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-192h-896q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm-896 448h896q26 0 45 19t19 45.5t-19 45t-45 18.5h-896q-27 0-45.5-18.5T.865 640t18.5-45.5t45.5-18.5zm0 192h384q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5T.865 832t18.5-45.5t45.5-18.5z"/>`),
		g.Group(children),
	)
}