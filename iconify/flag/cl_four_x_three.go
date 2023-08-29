package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCl4x30"><path fill-opacity=".7" d="M0 0h682.7v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagCl4x30)" transform="scale(.9375)"><path fill="#fff" d="M256 0h512v256H256z"/><path fill="#0039a6" d="M0 0h256v256H0z"/><path fill="#fff" d="M167.8 191.7L128.2 162l-39.5 30l14.7-48.8L64 113.1l48.7-.5L127.8 64l15.5 48.5l48.7.1l-39.2 30.4l15 48.7z"/><path fill="#d52b1e" d="M0 256h768v256H0z"/></g>`),
		g.Group(children),
	)
}