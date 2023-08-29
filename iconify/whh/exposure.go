package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exposure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-448-768h-64v-64q0-27-19-45.5t-45.5-18.5t-45 18.5t-18.5 45.5v64h-64q-27 0-45.5 19t-18.5 45.5t18.5 45t45.5 18.5h64v64q0 26 18.5 45t45 19t45.5-19t19-45v-64h64q26 0 45-18.5t19-45t-19-45.5t-45-19zm448 0q0-53-37-91l-694 694q38 37 91 37h512q53 0 90.5-37.5t37.5-90.5V256zm-128 512h-256q-27 0-45.5-19t-18.5-45.5t18.5-45t45.5-18.5h256q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}