package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><path id="cryptocurrencyColorOst0" d="m15.937 10.222l2.01-3.136a2.03 2.03 0 0 0 .83-3.97a2.03 2.03 0 0 0-2.04 3.2l-2.494 3.891a7.96 7.96 0 0 0-7.142 8.707a7.965 7.965 0 0 0 8.706 7.142a7.965 7.965 0 0 0 7.143-8.706a7.967 7.967 0 0 0-7.013-7.128zm-.915 12.122a4.22 4.22 0 0 1-4.224-4.223a4.22 4.22 0 0 1 4.224-4.224a4.22 4.22 0 0 1 3.95 2.732a2.442 2.442 0 0 0-3.424-.433a2.441 2.441 0 0 0-.433 3.424a2.441 2.441 0 0 0 3.849.007a4.214 4.214 0 0 1-3.942 2.717z"/></defs><g fill="none"><circle cx="16" cy="16" r="16" fill="#34445B"/><g fill="#FFF"><use href="#cryptocurrencyColorOst0"/><use href="#cryptocurrencyColorOst0"/></g></g>`),
		g.Group(children),
	)
}