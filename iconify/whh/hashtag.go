package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640q26 0 45 18.5t19 45.5v64q0 26-19 45t-45 19h-82l18 128q0 26-18.5 45t-45.5 19h-64q-26 0-45-19t-19-45l-18-128H430l18 128q0 26-19 45t-45 19h-64q-27 0-45.5-19T256 960l-18-128H64q-27 0-45.5-19T0 768v-64q0-27 18.5-45.5T64 640h146l-36-256H64q-27 0-45.5-19T0 320v-64q0-27 18.5-45.5T64 192h82L128 64q0-27 18.5-45.5T192 0h64q26 0 45 18.5T320 64l18 128h256L576 64q0-27 19-45.5T640 0h64q27 0 45.5 18.5T768 64l18 128h174q26 0 45 18.5t19 45.5v64q0 26-19 45t-45 19H814l36 256h110zM622 384H366l36 256h256z"/>`),
		g.Group(children),
	)
}