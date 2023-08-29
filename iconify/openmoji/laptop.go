package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M63.738 40.98v27.066H11.812V35.958h51.926zm-1.113-5.022h-50V3.891h50z"/><path fill="#FFF" d="m62.625 36l-50-22.221V3.892l50 22.221z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M62.625 40.98v27.066h-50V40.98m50-5.022h-50V3.891h50zm-45 5.57h40m-40 5.569h40m-40 5.569h40"/><path d="M30.717 56.186h13.506v8.104H30.717zm31.908-20.228h-50V3.891h50z"/></g>`),
		g.Group(children),
	)
}