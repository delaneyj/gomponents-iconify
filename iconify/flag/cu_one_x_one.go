package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CuOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCu1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagCu1x10)"><path fill="#002a8f" d="M-32 0h768v512H-32z"/><path fill="#fff" d="M-32 102.4h768v102.4H-32zm0 204.8h768v102.4H-32z"/><path fill="#cb1515" d="m-32 0l440.7 255.7L-32 511V0z"/><path fill="#fff" d="M161.8 325.5L114.3 290l-47.2 35.8l17.6-58.1l-47.2-36l58.3-.4l18.1-58l18.5 57.8l58.3.1l-46.9 36.3l18 58z"/></g>`),
		g.Group(children),
	)
}