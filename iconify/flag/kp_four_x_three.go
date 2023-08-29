package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KpFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKp4x30"><path fill-opacity=".7" d="M5 .1h682.6V512H5.1z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagKp4x30)" transform="translate(-4.8 -.1) scale(.93768)"><path fill="#fff" stroke="#000" d="M776 511.5H-76V.5h852z"/><path fill="#3e5698" d="M776 419H-76v92.5h852z"/><path fill="#c60000" d="M776 397.6H-76V114.4h852z"/><path fill="#3e5698" d="M776 .6H-76V93h852z"/><path fill="#fff" d="M328.5 256c0 63.5-53 115-118.6 115S91.3 319.5 91.3 256s53-114.8 118.6-114.8c65.5 0 118.6 51.4 118.6 114.9z"/><path fill="#c40000" d="m175.8 270.6l-57-40.7l71-.2l22.7-66.4l21.1 66.1l71-.4l-57.9 41.2l21.3 66.1l-57-40.7l-58 41.3z"/></g>`),
		g.Group(children),
	)
}