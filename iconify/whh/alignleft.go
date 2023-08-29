package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.865 320h-896q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm-192-192h-704q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h704q26 0 45 19t19 45.5t-18.5 45t-45.5 18.5zm-704 256h768q26 0 45 19t19 45.5t-19 45t-45 18.5h-768q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19zm0 192h832q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5h-832q-27 0-45.5-18.5T.865 640t18.5-45.5t45.5-18.5zm0 192h384q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5T.865 832t18.5-45.5t45.5-18.5z"/>`),
		g.Group(children),
	)
}