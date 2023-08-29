package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPrinterCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#fff" d="M7.5 5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v4.5h-11V5Z"/><path fill="#000" d="M8 9.5H7V5.1c0-.866.612-1.6 1.417-1.6h9.166c.805 0 1.417.734 1.417 1.6v4.4h-1V5.1c0-.35-.209-.6-.417-.6H8.417c-.208 0-.417.25-.417.6v4.4Z"/><path fill="#000" fill-rule="evenodd" d="M19 9H7a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3ZM5 12a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-5Z" clip-rule="evenodd"/><path fill="#fff" d="M18.5 20.531a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1V14.97h11v5.562Z"/><path fill="#000" d="M18 14.969h1v5.25c0 .97-.588 1.812-1.417 1.812H8.417C7.588 22.031 7 21.19 7 20.22v-5.25h1v5.25c0 .479.233.812.417.812h9.166c.184 0 .417-.333.417-.812v-5.25Z"/><path fill="#000" d="M16.5 18.5a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1h7Zm0-2a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1h7Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPrinterCircleFilled0)"/></g>`),
		g.Group(children),
	)
}